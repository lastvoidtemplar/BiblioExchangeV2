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
	AuthorId     string `json:"author_id"`
	Fullname     string `json:"fullname"`
	CountBooks   int    `json:"book_count"`
	CountStars   int    `json:"star_count"`
	CountReviews int    `json:"review_count"`
}

func MapAuthorsAToMinimizedAuthorDTOs(authors dbmodels.AuthorSlice) []MinimizedAuthorDTO {
	dtos := make([]MinimizedAuthorDTO, 0, len(authors))
	for _, v := range authors {
		log.Println(v.R.Authorpageratings)
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
	AuthorBooks   []MinimizedBookDTO  `json:"books"`
	AuthorReviews []*DisplayReviewDto `json:"reviews"`
	CountViews    int                 `json:"view_count"`
	CountStars    int                 `json:"star_count"`
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
	return SingleAuthorDTO{
		Author:        *author,
		AuthorBooks:   MapBooksAToMinimizedBookDTOs(author.R.Books),
		AuthorReviews: MapAuthorReviewSliceToDisplayReviewDtoTree(author.R.Authorreviews),
		CountStars:    countStars,
		CountViews:    countViews,
	}
}

type AuthorBodyDTO struct {
	Fullname     string `json:"fullname"`
	Biography    string `json:"biography"`
	DateOfBirth  string `json:"date_of_birth"`
	PlaceOfBirth string `json:"place_of_birth"`
	DateOfDeath  string `json:"date_of_death"`
	PlaceOfDeath string `json:"place_of_death"`
}

func (dto *AuthorBodyDTO) Valid() (errors validation.ValidationErrorsSlice) {
	if len(dto.Fullname) < 4 {
		errors.Add(fmt.Errorf("author fullname must be longer than 4 characters"))
	}

	if len(dto.Fullname) > 100 {
		errors.Add(fmt.Errorf("author fullname must be shorter than 100 characters"))
	}

	if dto.DateOfBirth != "" {
		if _, err := time.Parse(time.RFC3339, dto.DateOfBirth); err != nil {
			errors.Add(fmt.Errorf("author dateOfBirth must be in RFC-3339 / ISO-8601 format"))
		}
	}

	return
}

func (dto *AuthorBodyDTO) Map() (author dbmodels.Author) {
	author.Fullname = dto.Fullname
	author.Biography = dto.Biography

	if dto.DateOfBirth != "" {
		dateOfBirth, _ := time.Parse(time.RFC3339, dto.DateOfBirth)
		author.DateOfBirth = null.TimeFrom(dateOfBirth)
	}

	if dto.PlaceOfBirth != "" {
		author.PlaceOfBirth = null.StringFrom(dto.PlaceOfBirth)
	}

	if dto.DateOfDeath != "" {
		dateOfDeath, _ := time.Parse(time.RFC3339, dto.DateOfDeath)
		author.DateOfDeath = null.TimeFrom(dateOfDeath)
	}

	if dto.PlaceOfDeath != "" {
		author.PlaceOfDeath = null.StringFrom(dto.PlaceOfDeath)
	}

	return
}
