#!/bin/bash

set -e

echo "Waiting for PostgreSQL to be ready..."
max_retries=30
retry_interval=2

for ((i=1; i<=$max_retries; i++)); do
    if nc -z db 5432; then
        echo "PostgreSQL is ready!"

        # Дополнительная проверка, что база действительно принимает соединения
        if goose -dir ./internal/db/migrations postgres "user=postgres password=postgres dbname=library host=db port=5432 sslmode=disable" status >/dev/null 2>&1; then
            echo "Applying migrations..."
            goose -dir ./internal/db/migrations postgres "user=postgres password=postgres dbname=library host=db port=5432 sslmode=disable" up

            echo "Starting application..."
            exec ./gocionics
            exit 0
        else
            echo "PostgreSQL is listening but not yet accepting connections, retrying..."
        fi
    else
        echo "Attempt $i/$max_retries: PostgreSQL not ready, retrying in $retry_interval seconds..."
    fi
    sleep $retry_interval
done

echo "Failed to connect to PostgreSQL after $max_retries attempts"
exit 1