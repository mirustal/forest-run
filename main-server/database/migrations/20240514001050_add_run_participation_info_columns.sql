-- +goose Up
-- +goose StatementBegin
ALTER TABLE runs_participations
    ADD avg_tempo INTEGER DEFAULT 0,
    ADD created_at TIMESTAMP DEFAULT NOW(),
    ADD finish_time INTERVAL DEFAULT '00:00:00';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE runs_participations
    DROP COLUMN avg_tempo,
    DROP COLUMN created_at,
    DROP COLUMN finish_time;
-- +goose StatementEnd
