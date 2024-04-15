-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS runs_invites
(
    user_id INT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    run_id  INT NOT NULL REFERENCES runs (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS runs_invites;
-- +goose StatementEnd
