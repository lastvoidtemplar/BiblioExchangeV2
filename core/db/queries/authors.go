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
	authorPageRatingView = 0
	authorPageRatingLike = 1
)

func GetAllAuthors(ctx context.Context, db *sql.DB, pageNum, pageSize int) (dbmodels.AuthorSlice, error) {

	authors, err := dbmodels.Authors(
		LeftOuterJoin("authorpageratings r on r.author_id = author.author_id"),
		Where("rating_type = ?", authorPageRatingView),
		Or("r.author_id is null"),
		GroupBy("author.author_id"),
		OrderBy("count(author.author_id) desc"),
		Offset((pageNum-1)*pageSize),
		Limit(pageSize),
	).All(ctx, db)

	if err != nil {
		return nil, err
	}

	return authors, nil
}

func GetAuthorById(ctx context.Context, db *sql.DB, id string) (*dbmodels.Author, error) {

	author, err := dbmodels.Authors(
		Where("author.author_id = ?", id),
		Load(dbmodels.AuthorRels.Authorpageratings, Where("rating_type = ?", authorPageRatingLike)),
	).One(ctx, db)

	if err != nil {
		return nil, err
	}

	return author, nil
}

func AddAuthorPageView(ctx context.Context, db *sql.DB, authorId string, userId string, anonymous bool) error {
	var authorPageRating dbmodels.Authorpagerating
	authorPageRating.AuthorID = null.StringFrom(authorId)
	authorPageRating.RatingType = null.IntFrom(authorPageRatingView)
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
