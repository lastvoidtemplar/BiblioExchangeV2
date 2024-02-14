package dto

import (
	"fmt"

	dbmodels "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/models"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/server/validation"
	"github.com/volatiletech/null/v8"
)

type ListingBodyDto struct {
	Title         string   `json:"title"`
	Prize         int      `json:"prize"`
	Currency      string   `json:"currency"`
	Descrption    string   `json:"desc"`
	NumberOfFiles uint     `json:"files"`
	BooksIds      []string `json:"books"`
}

func (dto *ListingBodyDto) Valid() (errors validation.ValidationErrorsSlice) {
	if len(dto.Title) < 4 {
		errors.Add(fmt.Errorf("listing title must be longer than 4 characters"))
	}

	if len(dto.Title) > 100 {
		errors.Add(fmt.Errorf("listing title must be shorter than 100 characters"))
	}

	if len(dto.Descrption) > 5000 {
		errors.Add(fmt.Errorf("listing descrption must be shorter than 5000 characters"))
	}

	if dto.NumberOfFiles > 6 {
		errors.Add(fmt.Errorf("listing number of files must be less than 6"))
	}
	return
}

func (dto *ListingBodyDto) Map() (listings dbmodels.Listing) {
	listings.Title = dto.Title
	listings.Description = null.StringFrom(dto.Descrption)
	listings.Price = null.IntFrom(dto.Prize)
	listings.Currency = null.StringFrom(dto.Currency)

	return
}

type ListingSDisplayDto struct {
	ListingId  string
	Title      string
	Prize      int
	Currency   string
	Descrption string
	ImageUrls  []string
}

func MapListingsAToListingSDisplayDtos(listings dbmodels.ListingSlice) []ListingSDisplayDto {
	dtos := make([]ListingSDisplayDto, 0, len(listings))

	for _, l := range listings {
		dto := ListingSDisplayDto{
			ListingId:  l.ListingID,
			Title:      l.Title,
			Prize:      l.Price.Int,
			Currency:   l.Currency.String,
			Descrption: l.Description.String,
		}
		if l.R != nil {
			dto.ImageUrls = make([]string, 0, len(l.R.Listingsurls))
			for _, url := range l.R.Listingsurls {
				dto.ImageUrls = append(dto.ImageUrls, url.ResourseURL)
			}
		}
		dtos = append(dtos, dto)
	}
	return dtos
}
