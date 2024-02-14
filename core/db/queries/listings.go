package queries

import (
	"context"

	dbmodels "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func CreateListings(ctx context.Context, exec boil.ContextExecutor,
	listing *dbmodels.Listing, bookIds []string, imageNumber uint) ([]string, error) {
	err := listing.Insert(ctx, exec, boil.Blacklist(dbmodels.ListingColumns.ListingID))
	if err != nil {
		return nil, err
	}
	for _, id := range bookIds {
		book, err := dbmodels.FindBook(ctx, exec, id, dbmodels.BookColumns.BookID)
		if err != nil {
			return nil, err
		}
		err = listing.AddBooks(ctx, exec, false, book)
		if err != nil {
			return nil, err
		}
	}

	imagesUrl := make([]string, 0, imageNumber)
	for i := 0; i < int(imageNumber); i++ {
		img := &dbmodels.Listingsurl{
			ListingID: null.StringFrom(listing.ListingID),
		}
		err := img.Insert(ctx, exec, boil.Blacklist(dbmodels.ListingsurlColumns.ResourseID))
		if err != nil {
			return nil, err
		}
		imagesUrl = append(imagesUrl, img.ResourseID)
	}

	return imagesUrl, nil

}

func GetAllListings(ctx context.Context, db boil.ContextExecutor, bookId string, pageNum, pageSize int) (
	dbmodels.ListingSlice, error) {

	return dbmodels.Listings(
		InnerJoin("listingsbooks lb on lb.book_id = ?", bookId),
		Offset((pageNum-1)*pageSize),
		Limit(pageSize),
		Load(dbmodels.ListingRels.Listingsurls, Where("upload_status = ?", true)),
	).All(ctx, db)
}
