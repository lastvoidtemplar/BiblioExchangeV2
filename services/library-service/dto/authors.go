package dto

import (
	"fmt"
	"log"
	"time"

	dbmodels "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/models"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/db/queries"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/server/validation"
	"github.com/volatiletech/null/v8"
)

type MinimizedAuthorDTO struct {
	AuthorId     string
	Fullname     string
	CountBooks   int
	CountStars   int
	CountReviews int
}

func MapAuthorsAToMinimizedAuthorDTOs(authors dbmodels.AuthorSlice) []MinimizedAuthorDTO {
	dtos := make([]MinimizedAuthorDTO, 0, len(authors))
	for _, v := range authors {
		dtos = append(dtos, MinimizedAuthorDTO{
			AuthorId:     v.AuthorID,
			Fullname:     v.Fullname,
			CountBooks:   len(v.R.Books),
			CountStars:   len(v.R.Authorpageratings), //it fetches only ratings with stars
			CountReviews: len(v.R.Authorreviews),
		})
	}
	return dtos
}

type SingleAuthorDTO struct {
	dbmodels.Author
	AuthorBooks []MinimizedBookDTO
	CountViews  int
	CountStars  int
}

func MapAuthorAToSingleAuthorDTO(author *dbmodels.Author) SingleAuthorDTO {
	countStars, countViews := 0, 0

	if author.R != nil {
		for _, v := range author.R.Authorpageratings {
			if v.RatingType == null.IntFrom(queries.AuthorPageRatingView) {
				countViews++
			} else if v.RatingType == null.IntFrom(queries.AuthorPageRatingStar) {
				countStars++
			}
		}
	}
	log.Println(len(author.R.Books))
	return SingleAuthorDTO{
		Author:      *author,
		AuthorBooks: MapBooksAToMinimizedBookDTOs(author.R.Books),
		CountStars:  countStars,
		CountViews:  countViews,
	}
}

type AuthorBodyDTO struct {
	Fullname     string     `json:"fullname"`
	Biography    string     `json:"biography"`
	DateOfBirth  *time.Time `json:"date_of_birth"`
	PlaceOfBirth string     `json:"place_of_birth"`
	DateOfDeath  *time.Time `json:"date_of_death"`
	PlaceOfDeath string     `json:"place_of_death"`
}

func (dto *AuthorBodyDTO) Valid() (errors validation.ValidationErrorsSlice) {
	if len(dto.Fullname) < 4 {
		errors.Add(fmt.Errorf("author fullname must be longer than 4 characters"))
	}

	if len(dto.Fullname) > 100 {
		errors.Add(fmt.Errorf("author fullname must be shorter than 100 characters"))
	}

	return
}

func (dto *AuthorBodyDTO) Map() (author dbmodels.Author) {
	author.Fullname = dto.Fullname
	author.Biography = dto.Biography

	if dto.DateOfBirth != nil {
		author.DateOfBirth = null.TimeFrom(*dto.DateOfBirth)
	}

	if dto.PlaceOfBirth != "" {
		author.PlaceOfBirth = null.StringFrom(dto.PlaceOfBirth)
	}

	if dto.DateOfDeath != nil {
		author.DateOfDeath = null.TimeFrom(*dto.DateOfDeath)
	}

	if dto.PlaceOfDeath != "" {
		author.PlaceOfDeath = null.StringFrom(dto.PlaceOfDeath)
	}

	return
}
