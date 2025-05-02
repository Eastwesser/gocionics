FROM golang:1.23-alpine AS builder
WORKDIR /app
RUN apk add --no-cache make git

# Установка зависимостей
RUN go install github.com/pressly/goose/v3/cmd/goose@latest && \
    go install github.com/swaggo/swag/cmd/swag@latest

# Копируем только файлы модуля сначала
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код
COPY . .

# Генерируем документацию с явным указанием модуля
RUN swag init -g cmd/api/v1/main.go --output docs/swagger --parseDependency --parseInternal --dir ./

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o gocionics ./cmd/api/v1

FROM alpine:3.19
WORKDIR /app
RUN apk add --no-cache postgresql-client bash
COPY --from=builder /app/gocionics .
COPY --from=builder /app/internal/db/migrations ./internal/db/migrations
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY .env .
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

CMD ["/entrypoint.sh"]