package queries

import (
	"context"
	"database/sql"

	dbmodels "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetAllAuthors(ctx context.Context, db *sql.DB, pageNum, pageSize int) (dbmodels.AuthorSlice, error) {

	authors, err := dbmodels.Authors(
		InnerJoin("authorpageratings r on r.author_id = author.author_id"),
		Where("rating_type = 0"),
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
		Load(dbmodels.AuthorRels.Authorpageratings, Where("rating_type = 1")),
	).One(ctx, db)

	if err != nil {
		return nil, err
	}

	return author, nil
}
