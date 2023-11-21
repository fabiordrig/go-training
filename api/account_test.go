package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	mock_db "go-training/db/mock"
	db "go-training/db/sqlc"
	"go-training/utils"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccountApi(t *testing.T) {
	account := randomAccount()

	testCases := []struct {
		name         string
		accountID    int64
		buildStubs   func(store *mock_db.MockStore)
		checkStatus  func(t *testing.T, recorder *httptest.ResponseRecorder)
		checkAccount func(t *testing.T, recorder *httptest.ResponseRecorder, gotAccount db.Account)
	}{
		{
			name:      "OK",
			accountID: int64(account.ID),
			buildStubs: func(store *mock_db.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)
			},
			checkStatus: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
			checkAccount: func(t *testing.T, recorder *httptest.ResponseRecorder, gotAccount db.Account) {
				requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name:      "NotFound",
			accountID: int64(account.ID),
			buildStubs: func(store *mock_db.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(db.Account{}, sql.ErrNoRows)
			},
			checkStatus: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
			checkAccount: func(t *testing.T, recorder *httptest.ResponseRecorder, gotAccount db.Account) {
				requireBodyMatchAccount(t, recorder.Body, db.Account{})
			},
		},
		{
			name:      "InternalError",
			accountID: int64(account.ID),
			buildStubs: func(store *mock_db.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(db.Account{}, sql.ErrConnDone)
			},
			checkStatus: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
			checkAccount: func(t *testing.T, recorder *httptest.ResponseRecorder, gotAccount db.Account) {
				requireBodyMatchAccount(t, recorder.Body, db.Account{})
			},
		},
		{
			name:      "InvalidID",
			accountID: 0,
			buildStubs: func(store *mock_db.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Times(0)
			},
			checkStatus: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
			checkAccount: func(t *testing.T, recorder *httptest.ResponseRecorder, gotAccount db.Account) {
				requireBodyMatchAccount(t, recorder.Body, db.Account{})
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock_db.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/accounts/%d", tc.accountID)

			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkStatus(t, recorder)
			tc.checkAccount(t, recorder, account)

		})

	}

}

func randomAccount() db.Account {
	return db.Account{
		ID:       int32(utils.RandomInt(1, 1000)),
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)
}
