-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS UploadFiles (
    id  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    presignedURL VARCHAR,
    file_id UUID,
    allowedFileFormats VARCHAR [],
    maxSize INTEGER,
    user_id UUID,
    date_of_expration DATE,
    callbackAdrr VARCHAR
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS UploadFiles;
-- +goose StatementEnd