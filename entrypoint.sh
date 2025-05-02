#!/bin/sh

set -e

# Применяем миграции
echo "Applying database migrations..."
goose -dir ./migrations postgres "user=postgres password=postgres dbname=library host=db port=5432 sslmode=disable" up

# Запускаем приложение
echo "Starting application..."
exec ./gocionics