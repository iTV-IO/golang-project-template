-include .env

.PHONY: build tidy run debug

DB_URL=postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable
MIGRATE := migrate -path ./migrations -database $(DB_URL)
MIGRATE_NAME := $(if $(name),$(name),default_name)

build:
	@mkdir -p build
	@go build -o build/main cmd/main.go

tidy:
	@go mod tidy
	@go mod vendor

run:
	@GIN_MODE=release go run cmd/main.go

debug:
	@GIN_MODE=debug go run cmd/main.go

migrate-up:
	@$(MIGRATE) up
	@echo "Migrations applied successfully."

migrate-down:
	@$(MIGRATE) down
	@echo "Migrations rolled back successfully."

create-migration:
	@$(if $(name),,echo "Error: migration name required"; exit 1)
	@migrate create -ext sql -dir ./migrations -seq $(MIGRATE_NAME)
	@echo "Migration '$(MIGRATE_NAME)' created successfully."

migrate-force:
	@$(MIGRATE) force $(version)
	@echo "Migration forced successfully."

migrate-status:
	@$(MIGRATE) version
