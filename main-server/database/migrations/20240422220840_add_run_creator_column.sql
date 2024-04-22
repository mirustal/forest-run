-- +goose Up
-- +goose StatementBegin
ALTER TABLE runs
    ADD COLUMN creator INT NOT NULL DEFAULT 1 REFERENCES users(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE runs
    DROP COLUMN creator;
-- +goose StatementEnd
