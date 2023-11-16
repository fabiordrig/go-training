-include .env.local

postgres:
	docker run --name $(CONTAINER_NAME) -p 5432:5432 -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASS) -d postgres:alpine

createdb:
	@echo "Creating database..."
	@docker exec -it $(CONTAINER_NAME) createdb --username=$(DB_USER) --owner=$(DB_PASS) $(DB_NAME)
	@echo "Database created."

dropdb:
	@echo "Dropping database..."
	@docker exec -it $(CONTAINER_NAME) dropdb $(DB_NAME)
	@echo "Database dropped."

migrate:
	@echo "Migrating database...$(DB_NAME) $(DB_USER) $(DB_PASS)"
	migrate -path db/migrations -database "postgres://$(DB_USER):$(DB_PASS)@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose up

revert:
	migrate -path db/migrations -database "postgres://$(DB_USER):$(DB_PASS)@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrate revert sqlc test
