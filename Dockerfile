FROM golang:1.23-alpine AS builder
WORKDIR /app
RUN apk add --no-cache make git

# Критически важные переменные окружения
ENV GOPATH=/go
ENV GO111MODULE=on
ENV CGO_ENABLED=0

# Установка утилит
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Копирование и загрузка зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копирование исходного кода
COPY . .

# Генерация документации Swagger
RUN swag init -g cmd/api/v1/main.go --output docs/swagger

# Диагностика и сборка
RUN go list -m  # Показывает используемый модуль
RUN make build

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