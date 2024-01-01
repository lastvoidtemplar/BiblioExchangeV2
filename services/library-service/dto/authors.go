package dto

import (
	"fmt"
	"time"

	dbmodels "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/models"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/server/validation"
	"github.com/volatiletech/null/v8"
)

type SingleAuthorDTO struct {
	dbmodels.Author
	CountStars int
}

type AuthorCreateDTO struct {
	Fullname     string     `json:"fullname"`
	Biography    string     `json:"biography"`
	DateOfBirth  *time.Time `json:"date_of_birth"`
	PlaceOfBirth string     `json:"place_of_birth"`
	DateOfDeath  *time.Time `json:"date_of_death"`
	PlaceOfDeath string     `json:"place_of_death"`
}

func (dto *AuthorCreateDTO) Valid() (errors validation.ValidationErrorsSlice) {
	if len(dto.Fullname) < 4 {
		errors.Add(fmt.Errorf("author fullname must be longer than 4 characters"))
	}

	if len(dto.Fullname) > 100 {
		errors.Add(fmt.Errorf("author fullname must be shorter than 100 characters"))
	}

	return
}

func (dto *AuthorCreateDTO) Map() (author dbmodels.Author) {
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
