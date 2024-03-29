-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Listings (
    listing_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_id UUID NOT NULL,
    title VARCHAR(300) NOT NULL,
    price INTEGER,
    currency VARCHAR(10),
    description VARCHAR(5000)
);

CREATE TABLE IF NOT EXISTS ListingsURLs (
  resourse_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  listing_id UUID REFERENCES Listings(listing_id)  ON DELETE CASCADE,
  resourse_URL VARCHAR NOT NULL,
  upload_status BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS ListingsBooks (
    resourse_id UUID NOT NULL REFERENCES Listings(listing_id) ON DELETE CASCADE,
    book_id UUID NOT NULL REFERENCES Books(book_id)  ON DELETE CASCADE,
    PRIMARY KEY(resourse_id, book_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ListingsBooks;
DROP TABLE IF EXISTS ListingsURLs;
DROP TABLE IF EXISTS Listings;
-- +goose StatementEnd
