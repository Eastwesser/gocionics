FROM golang:1.23-alpine AS builder
WORKDIR /app
RUN apk add --no-cache make git
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build

FROM alpine:3.19
WORKDIR /app
# Устанавливаем зависимости
RUN apk add --no-cache postgresql-client bash
# Копируем файлы
COPY --from=builder /app/gocionics .
COPY --from=builder /app/internal/db/migrations ./internal/db/migrations
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Только один CMD!
CMD ["/entrypoint.sh"]