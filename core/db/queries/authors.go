package queries

import (
	"context"
	"database/sql"
	"errors"
	"log"

	dbmodels "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const (
	AuthorPageRatingView = 0
	AuthorPageRatingStar = 1
)

func GetAllAuthors(ctx context.Context, db *sql.DB, pageNum, pageSize int) (dbmodels.AuthorSlice, error) {

	authors, err := dbmodels.Authors(
		LeftOuterJoin("authorpageratings r on r.author_id = author.author_id"),
		GroupBy("author.author_id"),
		OrderBy("count(author.author_id) desc"),
		Offset((pageNum-1)*pageSize),
		Limit(pageSize),
		Load(dbmodels.AuthorRels.Books),
		Load(dbmodels.AuthorRels.Authorpageratings, Where("rating_type = ?", AuthorPageRatingStar)),
		Load(dbmodels.AuthorRels.Authorreviews),
	).All(ctx, db)

	if err != nil {
		return nil, err
	}

	return authors, nil
}

func GetAuthorById(ctx context.Context, db boil.ContextExecutor, id string) (*dbmodels.Author, error) {

	author, err := dbmodels.Authors(
		Where("author.author_id = ?", id),
		Load(dbmodels.AuthorRels.Authorpageratings),
		Load(dbmodels.AuthorRels.Books),
		Load(dbmodels.AuthorRels.Authorreviews),
	).One(ctx, db)

	if err != nil {
		return nil, err
	}

	return author, nil
}

func AddAuthorPageView(ctx context.Context, db boil.ContextExecutor, authorId string, userId string, anonymous bool) error {
	var authorPageRating dbmodels.Authorpagerating
	authorPageRating.AuthorID = null.StringFrom(authorId)
	authorPageRating.RatingType = null.IntFrom(AuthorPageRatingView)
	if !anonymous {
		authorPageRating.UserID = null.StringFrom(userId)
		log.Println(userId)
		return authorPageRating.Insert(ctx, db,
			boil.Blacklist(dbmodels.AuthorpageratingColumns.AuthorRatingID))
	}
	return authorPageRating.Insert(ctx, db,
		boil.Blacklist(
			dbmodels.AuthorpageratingColumns.AuthorRatingID,
			dbmodels.AuthorpageratingColumns.UserID,
		))

}

func CreateAuthor(ctx context.Context, db boil.ContextExecutor, author *dbmodels.Author) error {
	err := author.Insert(ctx, db, boil.Blacklist(dbmodels.AuthorColumns.AuthorID))

	return err

}

func UpdateAuthor(ctx context.Context, db *sql.DB, author *dbmodels.Author) (*dbmodels.Author, bool, error) {
	found, err := dbmodels.AuthorExists(ctx, db, author.AuthorID)

	if err != nil {
		return nil, false, err
	}

	if !found {
		return nil, false, nil
	}

	_, err = author.Update(ctx, db, boil.Blacklist(dbmodels.AuthorColumns.AuthorID))
	if err != nil {
		return nil, true, err
	}

	newAuthor, err := GetAuthorById(ctx, db, author.AuthorID)
	return newAuthor, true, err
}

var ErrAuthorHasBooks = errors.New("the author has books. deleted them first")

func DeleteAuthor(ctx context.Context, db *sql.DB, authorId string) (found bool, err error) {
	tx, err := db.BeginTx(ctx, nil)
	defer tx.Rollback()
	if err != nil {
		return false, err
	}

	author, err := dbmodels.Authors(
		Select(dbmodels.AuthorColumns.AuthorID),
		Where("author.author_id = ?", authorId),
		Load(dbmodels.AuthorRels.Books),
	).One(ctx, tx)

	if err != nil {
		err = nil
		return
	}

	found = true

	if author.R == nil || len(author.R.Books) != 0 {
		err = ErrAuthorHasBooks
		return
	}

	_, err = author.Delete(ctx, tx)

	if err != nil {
		return
	}

	err = tx.Commit()
	return
}

func createAuthorStarRating(ctx context.Context, exec boil.ContextExecutor, authorId string, userId string) error {
	authorRating := dbmodels.Authorpagerating{
		AuthorID:   null.StringFrom(authorId),
		UserID:     null.StringFrom(userId),
		RatingType: null.IntFrom(AuthorPageRatingStar),
	}

	return authorRating.Insert(ctx, exec, boil.Infer())
}

func ToggleStarRatingOnAuthorPage(ctx context.Context, db *sql.DB, authorId string, userId string) (starred bool, authorFound bool, err error) {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return false, false, err
	}
	defer tx.Rollback()
	authorFound, err = dbmodels.AuthorExists(ctx, tx, authorId)

	if err != nil {
		return false, false, err
	}

	if !authorFound {
		return false, false, nil
	}

	authorFound = true

	rating, err := dbmodels.Authorpageratings(
		Select(dbmodels.AuthorpageratingTableColumns.AuthorRatingID),
		Where("user_id = ?", userId),
		And("author_id = ?", authorId),
		And("rating_type = ?", AuthorPageRatingStar),
	).One(ctx, tx)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}

	err = nil
	starred = rating == nil

	if starred {
		err = createAuthorStarRating(ctx, tx, authorId, userId)
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

func ExistAuthorReviewOnSpecificAuthor(ctx context.Context, exec boil.ContextExecutor,
	rootId string, authorId string) (bool, error) {
	count, err := dbmodels.Authorreviews(
		Where("author_reviews_id = ?", rootId),
		And("author_id = ?", authorId),
	).Count(ctx, exec)

	if err != nil {
		return false, err
	}
	return count != 0, nil
}
func CreateAuthorReview(ctx context.Context, exec boil.ContextExecutor,
	authorId string, user_id string, content string, rootId string) (
	reviewId string, authorFound bool, rootFound bool, err error) {

	rootFound = true

	authorFound, err = dbmodels.AuthorExists(ctx, exec, authorId)

	if err != nil || !authorFound {
		return
	}
	authorFound = true

	review := dbmodels.Authorreview{
		AuthorID: authorId,
		UserID:   user_id,
		Content:  null.StringFrom(content),
	}

	if rootId != "" {
		rootFound = false
		rootFound, err = dbmodels.AuthorreviewExists(ctx, exec, rootId)
		if err != nil || !rootFound {
			return
		}
		rootFound = true
		review.RootID = null.StringFrom(rootId)
	}

	err = review.Insert(ctx, exec, boil.Blacklist(dbmodels.AuthorreviewColumns.AuthorReviewsID))

	if err != nil {
		return
	}
	reviewId = review.AuthorReviewsID
	return
}

func GetAuthorReviewById(ctx context.Context, exec boil.ContextExecutor, reviewId string) (*dbmodels.Authorreview, error) {
	return dbmodels.Authorreviews(
		Where("author_reviews_id = ?", reviewId),
	).One(ctx, exec)
}
