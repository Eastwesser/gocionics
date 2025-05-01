FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o gocionics ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/gocionics .
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./gocionics"]
