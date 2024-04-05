-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS run_posts
(
    id         BIGSERIAL NOT NULL PRIMARY KEY,
    post_type  int4      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id    INTEGER   NOT NULL REFERENCES users (id),
    run_id     INTEGER   NOT NULL REFERENCES runs (id),
    latitude   FLOAT     NOT NULL,
    longitude  FLOAT     NOT NULL,
    body       json      NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if EXISTS run_posts;
-- +goose StatementEnd
