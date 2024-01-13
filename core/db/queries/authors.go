package queries

import (
	"context"
	"database/sql"
	"fmt"
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

func GetAuthorById(ctx context.Context, db *sql.DB, id string) (*dbmodels.Author, error) {

	author, err := dbmodels.Authors(
		Where("author.author_id = ?", id),
		Load(dbmodels.AuthorRels.Authorpageratings),
		Load(dbmodels.AuthorRels.Books),
	).One(ctx, db)

	if err != nil {
		return nil, err
	}

	return author, nil
}

func AddAuthorPageView(ctx context.Context, db *sql.DB, authorId string, userId string, anonymous bool) error {
	fmt.Println(authorId, "  ", userId)

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

func CreateAuthor(ctx context.Context, db *sql.DB, author *dbmodels.Author) error {
	err := author.Insert(ctx, db, boil.Blacklist(dbmodels.AuthorColumns.AuthorID))

	return err

}

func UpdateAuthor(ctx context.Context, db *sql.DB, author *dbmodels.Author) error {
	_, err := author.Update(ctx, db, boil.Blacklist(dbmodels.AuthorColumns.AuthorID))

	return err

}

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
		err = fmt.Errorf("the author with id (%s) have books", authorId)
		return
	}

	_, err = author.Delete(ctx, tx)

	if err != nil {
		return
	}

	err = tx.Commit()
	return
}
