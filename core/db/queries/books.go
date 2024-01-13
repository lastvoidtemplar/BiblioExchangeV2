package queries

import (
	"context"
	"database/sql"
	"log"

	dbmodels "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const (
	BookPageRatingView = 0
	BookPageRatingLike = 1
)

func GetAllBooks(ctx context.Context, db *sql.DB, pageNum, pageSize int) (dbmodels.BookSlice, error) {

	books, err := dbmodels.Books(
		LeftOuterJoin("bookpageratings r on r.book_id = book.book_id"),
		GroupBy("book.book_id"),
		OrderBy("count(book.book_id) desc"),
		Offset((pageNum-1)*pageSize),
		Limit(pageSize),
	).All(ctx, db)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func GetBookById(ctx context.Context, db *sql.DB, id string) (*dbmodels.Book, error) {

	book, err := dbmodels.Books(
		Where("book.book_id = ?", id),
		Load(dbmodels.BookRels.Bookpageratings, Where("rating_type = ?", BookPageRatingLike)),
	).One(ctx, db)

	if err != nil {
		return nil, err
	}

	return book, nil
}

func AddBookPageView(ctx context.Context, db *sql.DB, bookId string, userId string, anonymous bool) error {
	var bookPageRating dbmodels.Bookpagerating
	bookPageRating.BookID = null.StringFrom(bookId)
	bookPageRating.RatingType = null.IntFrom(BookPageRatingView)
	if !anonymous {
		bookPageRating.UserID = null.StringFrom(userId)
		log.Println(userId)
		return bookPageRating.Insert(ctx, db,
			boil.Blacklist(dbmodels.BookpageratingColumns.BookRatingID))
	}
	return bookPageRating.Insert(ctx, db,
		boil.Blacklist(
			dbmodels.BookpageratingColumns.BookRatingID,
			dbmodels.BookpageratingColumns.UserID,
		))

}

func CreateBook(ctx context.Context, db *sql.DB, book *dbmodels.Book) error {
	err := book.Insert(ctx, db, boil.Blacklist(dbmodels.BookColumns.BookID))

	return err

}

func UpdateBook(ctx context.Context, db *sql.DB, book *dbmodels.Book) error {
	_, err := book.Update(ctx, db, boil.Blacklist(dbmodels.BookColumns.BookID))

	return err

}

func DeleteBook(ctx context.Context, db *sql.DB, bookId string) (found bool, err error) {
	tx, err := db.BeginTx(ctx, nil)

	if err != nil {
		return false, err
	}

	book, err := dbmodels.Books(
		Select(dbmodels.BookColumns.BookID),
		Where("book.book_id = ?", bookId),
	).One(ctx, tx)

	if err != nil {
		err = nil
		tx.Rollback()
		return
	}

	found = true

	_, err = book.Delete(ctx, tx)

	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}
