-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    email VARCHAR(50),
    password VARCHAR(50)
);
CREATE TABLE IF NOT EXISTS posts(
    id SERIAL PRIMARY KEY,
    owner_id INTEGER REFERENCES  users (id),
    title VARCHAR(100),
    content TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users CASCADE;
DROP TABLE posts;
-- +goose StatementEnd
