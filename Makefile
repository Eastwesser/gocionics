BINARY_NAME=gocionics
GO_FILES=$(shell find . -type f -name '*.go')
MIGRATIONS_DIR=./internal/db/migrations

.PHONY: build run test lint migrate-up migrate-down

build:
	@echo "Generating Swagger docs..."
	@swag init -g cmd/api/v1/main.go --output docs/swagger
	@echo "Building binary..."
	@CGO_ENABLED=0 GOOS=linux go build -o ${BINARY_NAME} ./cmd/api/v1

run:
	@echo "Starting server..."
	@go run ./cmd/api/v1

run-dev:
	@echo "Starting dev server..."
	@export DB_HOST=localhost && go run ./cmd/api/v1

test:
	@echo "Running tests..."
	@go test -v ./...

lint:
	@echo "Running linter..."
	@golangci-lint run

migrate-up:
	@echo "Applying migrations..."
	@goose -dir ${MIGRATIONS_DIR} postgres "user=postgres password=postgres dbname=library host=db port=5432 sslmode=disable" up

migrate-down:
	@echo "Reverting migrations..."
	@goose -dir ${MIGRATIONS_DIR} postgres "user=postgres password=postgres dbname=library host=db port=5432 sslmode=disable" down

swagger:
	@echo "Generating Swagger docs..."
	@swag init -g cmd/api/v1/main.go --output docs/swagger

deps:
	@echo "Installing dependencies..."
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

docker-build:
	@echo "Building Docker image..."
	@docker build -t ${BINARY_NAME} .

docker-run:
	@echo "Starting Docker container..."
	@docker run --rm -p 8080:8080 --network=host ${BINARY_NAME}

clean:
	@echo "Cleaning up..."
	@rm -f ${BINARY_NAME}