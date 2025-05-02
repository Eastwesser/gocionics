-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id            SERIAL PRIMARY KEY,
    email         TEXT UNIQUE NOT NULL,
    password_hash TEXT        NOT NULL,
    created_at    TIMESTAMP DEFAULT NOW(),
    updated_at    TIMESTAMP DEFAULT NOW()
);

CREATE TABLE characters
(
    id          SERIAL PRIMARY KEY,
    type        VARCHAR(50) NOT NULL,
    description TEXT,
    traits      TEXT[]
);

CREATE TABLE user_characters
(
    user_id      INT REFERENCES users (id),
    character_id INT REFERENCES characters (id),
    PRIMARY KEY (user_id, character_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_characters;
DROP TABLE IF EXISTS characters;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
