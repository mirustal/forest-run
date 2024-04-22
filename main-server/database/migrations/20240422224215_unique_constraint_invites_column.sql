-- +goose Up
-- +goose StatementBegin
ALTER TABLE runs_invites
    ADD CONSTRAINT unique_constraint_invites_column UNIQUE (user_id, run_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE runs_invites
    DROP CONSTRAINT unique_constraint_invites_column;
-- +goose StatementEnd
