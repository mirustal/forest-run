-- +goose Up
-- +goose StatementBegin
ALTER TABLE runs_participations
    ADD CONSTRAINT unique_constraint_participations_column UNIQUE (participant_id, run_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE runs_participations
    DROP CONSTRAINT unique_constraint_participations_column;
-- +goose StatementEnd
