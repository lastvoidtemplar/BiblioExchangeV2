-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Books (
    book_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ISBN VARCHAR(10) UNIQUE,
    title VARCHAR(100)  NOT NULL,
    date_of_publication DATE,
    plot VARCHAR(5000),
    genre VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS BookPageRatings (
    book_rating_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    book_id UUID REFERENCES Books(book_id)  ON DELETE CASCADE,
    rating_type INTEGER,
    user_id UUID
);

CREATE TABLE IF NOT EXISTS BookReviews (
    book_reviews_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    root_id UUID REFERENCES BookReviews(book_reviews_id)  ON DELETE CASCADE,
    book_id UUID REFERENCES Books(book_id)  ON DELETE CASCADE,
    user_id UUID  NOT NULL,
    content VARCHAR(1000)  NOT NULL
);

CREATE TABLE IF NOT EXISTS AuthorsBooks (
    author_id UUID NOT NULL REFERENCES Author(author_id)  ON DELETE CASCADE,
    book_id UUID NOT NULL REFERENCES Books(book_id)  ON DELETE CASCADE,
    PRIMARY KEY(author_id, book_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS AuthorsBooks;
DROP TABLE IF EXISTS BookReviews;
DROP TABLE IF EXISTS BookPageRatings;
DROP TABLE IF EXISTS Books;
-- +goose StatementEnd