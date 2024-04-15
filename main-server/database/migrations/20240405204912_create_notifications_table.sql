-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS notifications
(
    from_user_id INT  NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    to_user_id   INT  NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    type         INT  NOT NULL,
    created_at   TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    status       int4 NOT NULL DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notifications;
-- +goose StatementEnd
