MIGRATIONS_PATH=./cmd/migrate/migrations

# This pattern rule is necessary to handle arguments with filter-out
%:
	@:

.PHONY: migrate-create
migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out migrate-create,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) --database=$(DB_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate --path=$(MIGRATIONS_PATH) --database=$(DB_ADDR) down $(filter-out migrate-down,$(MAKECMDGOALS))

.PHONY: db-seed
db-seed:
	@go run ./cmd/migrate/seed/main.go

.PHONY: gen-docs
gen-docs:
	@swag init -g ./api/main.go -d cmd,internal && swag fmt
