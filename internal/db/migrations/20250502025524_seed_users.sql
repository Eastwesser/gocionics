-- +goose Up
-- +goose StatementBegin
DO
$$
    DECLARE
        user1_id       INT;
        user2_id       INT;
        don_quixote_id INT;
        duma_id        INT;
    BEGIN
        -- Вставляем пользователей
        INSERT INTO users (email, password_hash)
        VALUES ('user1@example.com', '$2a$10$XGVL6rMZQ7S6U6Zz7qJQY.FYgZ7hDf6QJQY8X6Xz7qJQY.FYgZ7hDf')
        ON CONFLICT (email) DO NOTHING
        RETURNING id INTO user1_id;

        INSERT INTO users (email, password_hash)
        VALUES ('user2@example.com', '$2a$10$XGVL6rMZQ7S6U6Zz7qJQY.FYgZ7hDf6QJQY8X6Xz7qJQY.FYgZ7hDf')
        ON CONFLICT (email) DO NOTHING
        RETURNING id INTO user2_id;

        -- Получаем ID персонажей
        SELECT id INTO don_quixote_id FROM characters WHERE type = 'Дон Кихот';
        SELECT id INTO duma_id FROM characters WHERE type = 'Дюма';

        -- Добавляем связи
        IF user1_id IS NOT NULL AND don_quixote_id IS NOT NULL THEN
            INSERT INTO user_characters (user_id, character_id)
            VALUES (user1_id, don_quixote_id)
            ON CONFLICT (user_id, character_id) DO NOTHING;
        END IF;

        IF user1_id IS NOT NULL AND duma_id IS NOT NULL THEN
            INSERT INTO user_characters (user_id, character_id)
            VALUES (user1_id, duma_id)
            ON CONFLICT (user_id, character_id) DO NOTHING;
        END IF;

        IF user2_id IS NOT NULL AND don_quixote_id IS NOT NULL THEN
            INSERT INTO user_characters (user_id, character_id)
            VALUES (user2_id, don_quixote_id)
            ON CONFLICT (user_id, character_id) DO NOTHING;
        END IF;
    END
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE users CASCADE;
-- +goose StatementEnd