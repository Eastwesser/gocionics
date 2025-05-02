-- +goose Up
-- +goose StatementBegin
-- Вставляем данные в таблицу characters (а не user_characters)
INSERT INTO characters (type, description, traits)
VALUES
    ('Дон Кихот', 'Искатель, интуитивно-логический экстраверт', ARRAY['изобретательный', 'энтузиаст', 'непредсказуемый']),
    ('Дюма', 'Посредник, сенсорно-этический интроверт', ARRAY['доброжелательный', 'гибкий', 'практичный']);

-- Пример добавления связи пользователь-характер (если нужно)
-- INSERT INTO user_characters (user_id, character_id)
-- VALUES (1, 1), (1, 2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Удаляем добавленные данные
TRUNCATE TABLE characters CASCADE;
-- +goose StatementEnd