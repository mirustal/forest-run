-- +goose Up
-- +goose StatementBegin
ALTER TABLE runs
    DROP COLUMN is_photo_allowed,
    DROP COLUMN is_stories_allowed,
    ADD COLUMN permissions_type INT NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE runs
    DROP COLUMN permissions_type,
    ADD COLUMN is_photo_allowed bool NOT NULL DEFAULT TRUE,
    ADD COLUMN is_stories_allowed bool NOT NULL DEFAULT TRUE;
-- +goose StatementEnd
