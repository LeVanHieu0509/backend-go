GOOSE_DBSTRING ?= "levanhieu:levanhieu1234@tcp(127.0.0.1:33060)/shopdevgo"
GOOSE_MIGRATION_DIR ?= sql/schema
GOOSE_DRIVER ?= mysql

APP_NAME = server


docker_build:
	docker-compose up -d --build
	docker-compose ps
docker_stop:
	docker-compose -f environment/docker-compose-dev.yml down
docker_up:
	docker-compose -f environment/docker-compose-dev.yml up -d
binancef:
	go run ./cmd/cli/binancef/main.binancef.go
dev: 
	go run ./cmd/${APP_NAME}/main.go

up_by_one: 
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up-by-one
# create migration
create_migration:
	@goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql

upse:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

downse:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

resetse:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

sqlgen: 
	sqlc generate
swag:
	swag init -g ./cmd/server/main.go -o ./cmd/swag/docs

.PHONY: dev downse upse resetse docker_build docker_stop docker_up swag

.PHONY: air

