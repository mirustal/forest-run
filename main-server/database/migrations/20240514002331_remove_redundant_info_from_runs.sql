-- +goose Up
-- +goose StatementBegin
ALTER TABLE runs
    DROP COLUMN permissions_type;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE runs
    ADD COLUMN permissions_type int4 DEFAULT 0;
-- +goose StatementEnd
