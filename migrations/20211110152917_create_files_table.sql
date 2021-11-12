-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS files
(
    id          bigserial PRIMARY KEY,
    hash        varchar(11) UNIQUE NOT NULL,
    language_id SERIAL NOT NULL,
    content     text NOT NULL,
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
ALTER TABLE files ADD CONSTRAINT file_languages_fkey FOREIGN KEY (language_id) REFERENCES languages(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS files;
-- +goose StatementEnd
