-- +goose Up
-- +goose StatementBegin
INSERT INTO public.characters (type, description, traits)
VALUES
    ('Дон Кихот', 'Искатель, интуитивно-логический экстраверт', ARRAY['изобретательный', 'энтузиаст', 'непредсказуемый']),
    ('Дюма', 'Посредник, сенсорно-этический интроверт', ARRAY['доброжелательный', 'гибкий', 'практичный']);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE public.characters CASCADE;
-- +goose StatementEnd