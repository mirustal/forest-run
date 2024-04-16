-- +goose Up
-- +goose StatementBegin
ALTER TABLE runs
    ADD COLUMN registration_until TIMESTAMP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE runs DROP COLUMN registration_until;
-- +goose StatementEnd
