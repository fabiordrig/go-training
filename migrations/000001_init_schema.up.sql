CREATE TABLE "accounts" (
  "id" serial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE "entries" (
  "id" serial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE "transfers" (
  "id" serial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now()
);

ALTER TABLE "entries" ADD CONSTRAINT "fk_entries_account"
  FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD CONSTRAINT "fk_transfers_from_account"
  FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD CONSTRAINT "fk_transfers_to_account"
  FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

CREATE INDEX "idx_accounts_owner" ON "accounts" ("owner");
CREATE INDEX "idx_entries_account_id" ON "entries" ("account_id");
CREATE INDEX "idx_transfers_from_account_id" ON "transfers" ("from_account_id");
CREATE INDEX "idx_transfers_to_account_id" ON "transfers" ("to_account_id");
CREATE INDEX "idx_transfers_from_to_account" ON "transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entries"."amount" IS 'Can be negative or positive';
COMMENT ON COLUMN "transfers"."amount" IS 'Must be positive';
