-- +goose Up
-- +goose StatementBegin
TRUNCATE TABLE public.user_characters CASCADE;
TRUNCATE TABLE public.characters CASCADE;
TRUNCATE TABLE public.users CASCADE;

INSERT INTO public.users (email, password_hash)
VALUES
    ('user1@example.com', '$2a$10$XGVL6rMZQ7S6U6Zz7qJQY.FYgZ7hDf6QJQY8X6Xz7qJQY.FYgZ7hDf'),
    ('user2@example.com', '$2a$10$XGVL6rMZQ7S6U6Zz7qJQY.FYgZ7hDf6QJQY8X6Xz7qJQY.FYgZ7hDf');

INSERT INTO public.characters (type, description, traits)
VALUES
    ('Дон Кихот', 'Искатель, интуитивно-логический экстраверт', ARRAY['изобретательный', 'энтузиаст', 'непредсказуемый']),
    ('Дюма', 'Посредник, сенсорно-этический интроверт', ARRAY['доброжелательный', 'гибкий', 'практичный']);

INSERT INTO public.user_characters (user_id, character_id)
VALUES
    (1, 1),
    (1, 2),
    (2, 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE public.user_characters CASCADE;
TRUNCATE TABLE public.characters CASCADE;
TRUNCATE TABLE public.users CASCADE;
-- +goose StatementEnd