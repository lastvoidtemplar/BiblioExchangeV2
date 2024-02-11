package queries

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	dbmodels "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const (
	BookPageRatingView = 0
	BookPageRatingStar = 1
)

func GetAllBooks(ctx context.Context, db boil.ContextExecutor, authorId string, pageNum, pageSize int) (dbmodels.BookSlice, error) {

	books, err := dbmodels.Books(
		InnerJoin("authorsbooks ab on ab.author_id = ?", authorId),
		LeftOuterJoin("bookpageratings r on r.book_id = books.book_id"),
		GroupBy("books.book_id"),
		OrderBy("count(books.book_id) desc"),
		Offset((pageNum-1)*pageSize),
		Limit(pageSize),
		Load(dbmodels.BookRels.Bookpageratings, Where("rating_type = ?", BookPageRatingStar)),
		Load(dbmodels.BookRels.Bookreviews),
	).All(ctx, db)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func GetBookById(ctx context.Context, db boil.ContextExecutor, id string) (*dbmodels.Book, error) {

	book, err := dbmodels.Books(
		Where("books.book_id = ?", id),
		Load(dbmodels.BookRels.Authors),
		Load(dbmodels.BookRels.Bookreviews),
		Load(dbmodels.BookRels.Bookpageratings),
	).One(ctx, db)

	if err != nil {
		return nil, err
	}
	if len(book.R.Authors) < 1 {
		return nil, fmt.Errorf("missing author")
	}
	book.R.Authors[0].L.LoadAuthorpageratings(ctx, db, true, book.R.Authors[0],
		Where("rating_type = ?", AuthorPageRatingStar))
	book.R.Authors[0].L.LoadAuthorreviews(ctx, db, true, book.R.Authors[0], nil)

	return book, nil
}

func AddBookPageView(ctx context.Context, db boil.ContextExecutor, bookId string, userId string, anonymous bool) error {
	var bookPageRating dbmodels.Bookpagerating
	bookPageRating.BookID = null.StringFrom(bookId)
	bookPageRating.RatingType = null.IntFrom(BookPageRatingView)
	if !anonymous {
		bookPageRating.UserID = null.StringFrom(userId)
		return bookPageRating.Insert(ctx, db,
			boil.Blacklist(dbmodels.BookpageratingColumns.BookRatingID))
	}
	return bookPageRating.Insert(ctx, db,
		boil.Blacklist(
			dbmodels.BookpageratingColumns.BookRatingID,
			dbmodels.BookpageratingColumns.UserID,
		))

}

func CreateBook(ctx context.Context, db boil.ContextExecutor, book *dbmodels.Book, authorId string) error {
	err := book.Insert(ctx, db, boil.Blacklist(dbmodels.BookColumns.BookID))
	if err != nil {
		return err
	}
	author, err := GetAuthorById(ctx, db, authorId)
	if err != nil {
		return err
	}
	return book.AddAuthors(ctx, db, false, author)
}

func UpdateBook(ctx context.Context, db boil.ContextExecutor, book *dbmodels.Book) (
	*dbmodels.Book, error) {
	_, err := book.Update(ctx, db, boil.Blacklist(dbmodels.BookColumns.BookID))

	if err != nil {
		return nil, err
	}

	return GetBookById(ctx, db, book.BookID)

}

func DeleteBook(ctx context.Context, db *sql.DB, bookId string) (found bool, err error) {
	tx, err := db.BeginTx(ctx, nil)

	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	book, err := dbmodels.Books(
		Select(dbmodels.BookColumns.BookID),
		Where("books.book_id = ?", bookId),
	).One(ctx, tx)

	if err != nil {
		err = nil
		return
	}

	found = true

	_, err = book.Delete(ctx, tx)

	if err != nil {
		return
	}

	err = tx.Commit()
	return
}

func createBookStarRating(ctx context.Context, exec boil.ContextExecutor, bookId string, userId string) error {
	authorRating := dbmodels.Bookpagerating{
		BookID:     null.StringFrom(bookId),
		UserID:     null.StringFrom(userId),
		RatingType: null.IntFrom(BookPageRatingStar),
	}

	return authorRating.Insert(ctx, exec, boil.Infer())
}

func ToggleStarRatingOnBookPage(ctx context.Context, db *sql.DB, bookId string, userId string) (
	starred bool, bookFound bool, err error) {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return false, false, err
	}
	defer tx.Rollback()
	bookFound, err = dbmodels.BookExists(ctx, tx, bookId)

	if err != nil {
		return false, false, err
	}

	if !bookFound {
		return false, false, nil
	}

	bookFound = true

	rating, err := dbmodels.Bookpageratings(
		Select(dbmodels.BookpageratingColumns.BookRatingID),
		Where("user_id = ?", userId),
		And("book_id = ?", bookId),
		And("rating_type = ?", BookPageRatingStar),
	).One(ctx, tx)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}

	err = nil
	starred = rating == nil

	if starred {
		err = createBookStarRating(ctx, tx, bookId, userId)
		if err == nil {
			tx.Commit()
		}
		return
	} else {
		_, err = rating.Delete(ctx, tx)
		if err == nil {
			tx.Commit()
		}
		return
	}
}

func ExistBookReviewOnSpecificBook(ctx context.Context, exec boil.ContextExecutor,
	rootId string, bookId string) (bool, error) {
	count, err := dbmodels.Bookreviews(
		Where("book_reviews_id = ?", rootId),
		And("book_id = ?", bookId),
	).Count(ctx, exec)

	if err != nil {
		return false, err
	}
	return count != 0, nil
}

func CreateBookReview(ctx context.Context, exec boil.ContextExecutor,
	bookId string, user_id string, content string, rootId string) (
	reviewId string, bookFound bool, rootFound bool, err error) {

	rootFound = true

	bookFound, err = dbmodels.BookExists(ctx, exec, bookId)

	if err != nil || !bookFound {
		return
	}
	bookFound = true

	review := dbmodels.Bookreview{
		BookID:  null.StringFrom(bookId),
		UserID:  user_id,
		Content: content,
	}

	if rootId != "" {
		rootFound = false
		rootFound, err = dbmodels.BookreviewExists(ctx, exec, rootId)
		if err != nil || !rootFound {
			return
		}
		rootFound = true
		review.RootID = null.StringFrom(rootId)
	}

	err = review.Insert(ctx, exec, boil.Blacklist(dbmodels.BookreviewColumns.BookReviewsID))

	if err != nil {
		return
	}
	reviewId = review.BookReviewsID
	return
}

func GetBookReviewById(ctx context.Context, exec boil.ContextExecutor, reviewId string) (*dbmodels.Bookreview, error) {
	return dbmodels.Bookreviews(
		Where("book_reviews_id = ?", reviewId),
	).One(ctx, exec)
}
