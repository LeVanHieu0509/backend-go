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

.PHONY: run

.PHONY: air