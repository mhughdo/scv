-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS languages
(
    id         SERIAL PRIMARY KEY,
    code       varchar(5) NOT NULL,
    name       varchar(15) NOT NULL,
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS languages;
-- +goose StatementEnd
