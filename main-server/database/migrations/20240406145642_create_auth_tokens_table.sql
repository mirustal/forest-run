-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS auth_tokens
(
    user_id    INTEGER   NOT NULL REFERENCES users (id),
    token      CHAR(256) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if EXISTS auth_tokens;
-- +goose StatementEnd
