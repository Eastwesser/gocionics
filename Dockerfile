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
# Добавляем make и зависимости для goose
RUN apk add --no-cache postgresql-client
COPY --from=builder /app/gocionics .
COPY --from=builder /app/internal/db/migrations ./migrations
COPY --from=builder /go/bin/goose /usr/local/bin/goose

# Заменяем make на прямой вызов goose и приложения
CMD ["sh", "-c", "goose -dir ./migrations postgres \"user=postgres password=postgres dbname=library host=db port=5432 sslmode=disable\" up && ./gocionics"]