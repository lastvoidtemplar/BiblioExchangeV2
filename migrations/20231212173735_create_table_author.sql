-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Author(
    author_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    fullname VARCHAR(100) NOT NUll,
    biography VARCHAR(5000) NOT NUll,
    date_of_birth DATE,
    place_of_birth VARCHAR(100),
    date_of_death DATE,
    place_of_death VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS AuthorPageRatings(
    author_rating_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    author_id UUID REFERENCES Author(author_id) ON DELETE CASCADE,
    rating_type INTEGER,
    user_id UUID
);

CREATE TABLE IF NOT EXISTS AuthorReviews (
    author_reviews_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    root_id UUID REFERENCES AuthorReviews(author_reviews_id)  ON DELETE CASCADE,
    author_id UUID NOT NULL REFERENCES Author(author_id) ON DELETE CASCADE,
    user_id UUID  NOT NULL,
    content VARCHAR(500)
); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS AuthorReviews;
DROP TABLE IF EXISTS AuthorPageRatings;
DROP TABLE IF EXISTS Author;
-- +goose StatementEnd