FROM golang:1.23-alpine AS builder

WORKDIR /app

# Установка зависимостей
RUN apk add --no-cache make git
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/gocionics .
COPY --from=builder /app/internal/db/migrations ./migrations
COPY --from=builder /go/bin/goose /usr/local/bin/goose

CMD ["sh", "-c", "make migrate-up && ./gocionics"]