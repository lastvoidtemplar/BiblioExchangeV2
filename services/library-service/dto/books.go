package dto

import (
	"fmt"
	"time"

	dbmodels "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/models"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/db/queries"
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
	Author       MinimizedAuthorDTO  `json:"author"`
	BooksReviews []*DisplayReviewDto `json:"reviews"`
	CountViews   int                 `json:"view_count"`
	CountStars   int                 `json:"star_count"`
}

func MapFromBook(book *dbmodels.Book) SingleBookDTO {
	countStars, countViews := 0, 0

	if book.R != nil {
		for _, v := range book.R.Bookpageratings {
			if v.RatingType == null.IntFrom(queries.BookPageRatingView) {
				countViews++
			} else if v.RatingType == null.IntFrom(queries.BookPageRatingStar) {
				countStars++
			}
		}
	}

	return SingleBookDTO{
		Book:         *book,
		Author:       MapAuthorsAToMinimizedAuthorDTOs(book.R.Authors[:1])[0],
		BooksReviews: MapBookReviewSliceToDisplayReviewDtoTree(book.R.Bookreviews),
		CountStars:   countStars,
		CountViews:   countViews,
	}
}

type BookBodyDTO struct {
	Isbn              string `json:"isbn"`
	Title             string `json:"title"`
	DateOfPublication string `json:"date_of_publication"`
	Plot              string `json:"plot"`
	Genre             string `json:"genre"`
}

func (dto *BookBodyDTO) Valid() (errors validation.ValidationErrorsSlice) {
	if len(dto.Isbn) != 10 {
		errors.Add(fmt.Errorf("book ISBN must be 10 characters long"))
	}

	if len(dto.Title) < 4 {
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

	if dto.DateOfPublication != "" {
		if _, err := time.Parse(time.RFC3339, dto.DateOfPublication); err != nil {
			errors.Add(fmt.Errorf("author date of publication must be in RFC-3339 / ISO-8601 format"))
		}
	}
	return
}

func (dto *BookBodyDTO) Map() (book dbmodels.Book) {
	book.Isbn = null.StringFrom(dto.Isbn)
	book.Title = dto.Title
	book.Plot = null.StringFrom(dto.Plot)

	if dto.DateOfPublication != "" {
		date, _ := time.Parse(time.RFC3339, dto.DateOfPublication)
		book.DateOfPublication = null.TimeFrom(date)
	}

	book.Genre = null.StringFrom(dto.Genre)

	return
}
