-- +goose Up
-- +goose StatementBegin
ALTER TABLE AuthorPageRatings
ADD COLUMN amount INTEGER;

ALTER TABLE BookPageRatings
ADD COLUMN amount INTEGER;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE AuthorPageRatings
DROP COLUMN amount;

ALTER TABLE BookPageRatings
DROP COLUMN amount;
-- +goose StatementEnd
