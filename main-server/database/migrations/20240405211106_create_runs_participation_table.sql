-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS runs_participations
(
    participant_id INT NOT NULL REFERENCES users (id),
    run_id         INT NOT NULL REFERENCES runs (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS runs_participations;
-- +goose StatementEnd
