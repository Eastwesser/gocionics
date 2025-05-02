#!/bin/bash

set -e

echo "Waiting for PostgreSQL to be ready..."
while ! nc -z db 5432; do
  sleep 1
done

echo "Applying migrations..."
goose -dir ./internal/db/migrations postgres "user=postgres password=postgres dbname=library host=db port=5432 sslmode=disable" up

echo "Starting application..."
exec ./gocionics