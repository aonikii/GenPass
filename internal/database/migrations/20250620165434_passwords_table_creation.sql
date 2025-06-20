-- +goose Up
-- +goose StatementBegin
CREATE TABLE passwords (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    site TEXT NOT NULL,
    password TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE passwords;
-- +goose StatementEnd
