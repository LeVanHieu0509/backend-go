GOOSE_DBSTRING = "levanhieu:levanhieu1234@tcp(127.0.0.1:33060)/shopdevgo"
GOOSE_MIGRATION_DIR ?= sql/schema

APP_NAME = server


binancef:
	go run ./cmd/cli/binancef/main.binancef.go
dev: 
	go run ./cmd/${APP_NAME}/main.go
run:
	docker compose up -d && go run ./cmd/${APP_NAME}/main.go
up:
	docker compose up -d
down:
	docker compose down

upse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

downse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

resetse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset


.PHONY: run downse upse resetse

.PHONY: air

