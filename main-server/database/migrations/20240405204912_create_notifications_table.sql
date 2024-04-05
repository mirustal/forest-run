-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS notifications
(
    user_id INT NOT NULL REFERENCES users (id),
    body       json   NOT NULL,
    created_at TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    status     int4   NOT NULL DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notifications;
-- +goose StatementEnd
