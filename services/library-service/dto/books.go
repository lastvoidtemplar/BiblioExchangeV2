package dto

import (
	"fmt"
	"time"

	dbmodels "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/models"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/server/validation"
	"github.com/volatiletech/null/v8"
)

type MinimizedBookDTO struct {
	BookId       string
	Title        string
	Genre        string
	CountStars   int
	CountReviews int
}

func MapBooksAToMinimizedBookDTOs(books dbmodels.BookSlice) []MinimizedBookDTO {
	dtos := make([]MinimizedBookDTO, 0, len(books))
	for _, v := range books {
		dtos = append(dtos, MinimizedBookDTO{
			BookId:       v.BookID,
			Title:        v.Title,
			Genre:        v.Genre.String,
			CountStars:   len(v.R.Bookpageratings), //it fetches only ratings with stars
			CountReviews: len(v.R.Bookreviews),
		})
	}
	return dtos
}

type SingleBookDTO struct {
	dbmodels.Book
	CountStars int
}

func MapFromBook(book *dbmodels.Book) SingleBookDTO {
	countStars := 0

	if book.R != nil {
		countStars = len(book.R.Bookpageratings)
	}

	return SingleBookDTO{
		Book:       *book,
		CountStars: countStars,
	}
}

type BookBodyDTO struct {
	Isbn              string     `json:"isbn"`
	Title             string     `json:"title"`
	DateOfPublication *time.Time `json:"date_of_publication"`
	Plot              string     `json:"plot"`
	Genre             string     `json:"genre"`
}

func (dto *BookBodyDTO) Valid() (errors validation.ValidationErrorsSlice) {
	if dto.Isbn != "" && len(dto.Isbn) != 10 {
		errors.Add(fmt.Errorf("book ISBN must be 10 characters long"))
	}

	if len(dto.Title) < 3 {
		errors.Add(fmt.Errorf("book title must be longer than 4 characters"))
	}

	if len(dto.Title) > 100 {
		errors.Add(fmt.Errorf("book title must be shorter than 100 characters"))
	}

	if len(dto.Genre) < 4 {
		errors.Add(fmt.Errorf("book genre must be longer than 4 characters"))
	}

	if len(dto.Genre) > 50 {
		errors.Add(fmt.Errorf("book genre must be shorter than 50 characters"))
	}
	return
}

func (dto *BookBodyDTO) Map() (book dbmodels.Book) {
	book.Isbn = null.StringFrom(dto.Isbn)
	book.Plot = null.StringFrom(dto.Plot)

	if dto.DateOfPublication != nil {
		book.DateOfPublication = null.TimeFrom(*dto.DateOfPublication)
	}

	book.Genre = null.StringFrom(dto.Genre)

	return
}
