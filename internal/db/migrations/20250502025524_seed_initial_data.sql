-- +goose Up
-- +goose StatementBegin

-- 1. Сначала очищаем таблицы если они уже существуют (для идемпотентности)
TRUNCATE TABLE user_characters CASCADE;
TRUNCATE TABLE characters CASCADE;
TRUNCATE TABLE users CASCADE;

-- 2. Вставляем базовых пользователей
INSERT INTO users (email, password_hash)
VALUES
    ('user1@example.com', '$2a$10$XGVL6rMZQ7S6U6Zz7qJQY.FYgZ7hDf6QJQY8X6Xz7qJQY.FYgZ7hDf'),
    ('user2@example.com', '$2a$10$XGVL6rMZQ7S6U6Zz7qJQY.FYgZ7hDf6QJQY8X6Xz7qJQY.FYgZ7hDf');

-- 3. Вставляем характеры
INSERT INTO characters (type, description, traits)
VALUES
    ('Дон Кихот', 'Искатель, интуитивно-логический экстраверт', ARRAY['изобретательный', 'энтузиаст', 'непредсказуемый']),
    ('Дюма', 'Посредник, сенсорно-этический интроверт', ARRAY['доброжелательный', 'гибкий', 'практичный']);

-- 4. Связываем пользователей с характерами
INSERT INTO user_characters (user_id, character_id)
VALUES
    (1, 1), -- user1@example.com -> Дон Кихот
    (1, 2), -- user1@example.com -> Дюма
    (2, 1); -- user2@example.com -> Дон Кихот

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Откат - полная очистка таблиц
TRUNCATE TABLE user_characters CASCADE;
TRUNCATE TABLE characters CASCADE;
TRUNCATE TABLE users CASCADE;
-- +goose StatementEnd