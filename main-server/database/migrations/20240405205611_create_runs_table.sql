-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS runs
(
    id                    SERIAL       NOT NULL PRIMARY KEY,
    name                  VARCHAR(128) NOT NULL,
    description           TEXT,
    official_site         TEXT,
    avatar_url            TEXT,
    route                 json         NOT NULL,
    start_time            TIMESTAMP    NOT NULL,
    start_place           TEXT         NOT NULL,
    start_place_latitude  FLOAT         NOT NULL,
    start_place_longitude FLOAT         NOT NULL,
    max_participants      INT          NOT NULL,
    is_photo_allowed      bool         NOT NULL DEFAULT TRUE,
    is_stories_allowed    bool         NOT NULL DEFAULT TRUE,
    status                int4         NOT NULL DEFAULT 0,
    participation_format  int4         NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS runs;
-- +goose StatementEnd
