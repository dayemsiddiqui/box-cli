-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS tasks (id TEXT PRIMARY KEY, name TEXT, description TEXT, created_at datetime, updated_at datetime);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tasks;
-- +goose StatementEnd
