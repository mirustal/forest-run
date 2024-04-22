-- +goose Up
-- +goose StatementBegin
ALTER TABLE notifications
    ADD COLUMN body json NOT NULL DEFAULT '{}';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE notifications
    DROP COLUMN body;
-- +goose StatementEnd
