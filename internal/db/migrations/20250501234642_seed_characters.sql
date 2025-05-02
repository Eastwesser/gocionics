-- +goose Up
-- +goose StatementBegin
DO
$$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM characters WHERE type = 'Дон Кихот') THEN
            INSERT INTO characters (type, description, traits)
            VALUES ('Дон Кихот',
                    'Искатель, интуитивно-логический экстраверт',
                    ARRAY ['изобретательный', 'энтузиаст', 'непредсказуемый']);
        END IF;

        IF NOT EXISTS (SELECT 1 FROM characters WHERE type = 'Дюма') THEN
            INSERT INTO characters (type, description, traits)
            VALUES ('Дюма',
                    'Посредник, сенсорно-этический интроверт',
                    ARRAY ['доброжелательный', 'гибкий', 'практичный']);
        END IF;
    END
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE characters CASCADE;
-- +goose StatementEnd