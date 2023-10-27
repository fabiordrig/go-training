-include local.env

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
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASS)@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose up

revert:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASS)@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose down


.PHONY: postgres createdb dropdb migrate revert
