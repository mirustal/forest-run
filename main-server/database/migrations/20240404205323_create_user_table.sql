-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id                       SERIAL PRIMARY KEY,
    username                    VARCHAR(64) NOT NULL UNIQUE,
    password                 CHAR(64)    NOT NULL,
    name                     VARCHAR(64),
    avatar_url               TEXT,
    refresh_token            CHAR(64),
    refresh_token_expires_at TIMESTAMP,
    created_at               TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at               TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
