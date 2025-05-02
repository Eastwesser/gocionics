FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o gocionics ./cmd/api/v1

FROM alpine:latest
WORKDIR /app

RUN apk add --no-cache tzdata

COPY --from=builder /app/gocionics .
COPY --from=builder /app/internal/db/migrations ./migrations
COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["sh", "-c", "goose -dir ./migrations postgres \"user=postgres password=postgres dbname=library host=db port=5432 sslmode=disable\" up && ./gocionics"]