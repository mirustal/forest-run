-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS subscriptions
(
    follower_id INT NOT NULL REFERENCES users (id),
    followed_id INT NOT NULL REFERENCES users (id),
    PRIMARY KEY (follower_id, followed_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS subscriptions;
-- +goose StatementEnd
