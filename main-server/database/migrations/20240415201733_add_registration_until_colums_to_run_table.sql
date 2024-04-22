-- +goose Up
-- +goose StatementBegin
ALTER TABLE runs
    ADD COLUMN registration_until TIMESTAMP NOT NULL DEFAULT '2100-01-01 00:00:00';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE runs DROP COLUMN registration_until;
-- +goose StatementEnd
