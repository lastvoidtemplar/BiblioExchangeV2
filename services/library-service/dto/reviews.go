package dto

import (
	"fmt"

	dbmodels "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/models"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/server/validation"
	"github.com/volatiletech/null/v8"
)

type ReviewBodyDto struct {
	Content string `json:"content"`
}

func (dto *ReviewBodyDto) Valid() (errors validation.ValidationErrorsSlice) {
	if len(dto.Content) < 3 {
		errors.Add(fmt.Errorf("author review content must be longer than 3 characters"))
	}
	if len(dto.Content) > 500 {
		errors.Add(fmt.Errorf("author review content must be shorter than 500 characters"))
	}
	return
}

type DisplayReviewDto struct {
	ReviewId string              `json:"review_id"`
	Content  string              `json:"content"`
	UserID   string              `json:"user_id"`
	Reviews  []*DisplayReviewDto `json:"reviews"`
}

func MapAuthorReviewToDisplayReviewDto(review dbmodels.Authorreview) DisplayReviewDto {
	return DisplayReviewDto{
		ReviewId: review.AuthorReviewsID,
		Content:  review.Content.String,
		UserID:   review.UserID,
		Reviews:  make([]*DisplayReviewDto, 0),
	}
}

func MapBookReviewToDisplayReviewDto(review dbmodels.Bookreview) DisplayReviewDto {
	return DisplayReviewDto{
		ReviewId: review.BookReviewsID,
		Content:  review.Content,
		UserID:   review.UserID,
		Reviews:  make([]*DisplayReviewDto, 0),
	}
}

func MapAuthorReviewSliceToDisplayReviewDtoTree(reviews dbmodels.AuthorreviewSlice) []*DisplayReviewDto {
	reviewsMap := make(map[string]*DisplayReviewDto)
	result := make([]*DisplayReviewDto, 0)
	for _, review := range reviews {
		temp := MapAuthorReviewToDisplayReviewDto(*review)
		reviewsMap[review.AuthorReviewsID] = &temp
	}
	var zeroValue null.String
	for _, review := range reviews {
		if review.RootID != zeroValue {

			reviewsMap[review.RootID.String].Reviews =
				append(reviewsMap[review.RootID.String].Reviews,
					reviewsMap[review.AuthorReviewsID])
		} else {
			result = append(result, reviewsMap[review.AuthorReviewsID])
		}
	}
	return result
}

func MapBookReviewSliceToDisplayReviewDtoTree(reviews dbmodels.BookreviewSlice) []*DisplayReviewDto {
	reviewsMap := make(map[string]*DisplayReviewDto)
	result := make([]*DisplayReviewDto, 0)
	for _, review := range reviews {
		temp := MapBookReviewToDisplayReviewDto(*review)
		reviewsMap[review.BookReviewsID] = &temp
	}
	var zeroValue null.String
	for _, review := range reviews {
		if review.RootID != zeroValue {

			reviewsMap[review.RootID.String].Reviews =
				append(reviewsMap[review.RootID.String].Reviews,
					reviewsMap[review.BookReviewsID])
		} else {
			result = append(result, reviewsMap[review.BookReviewsID])
		}
	}
	return result
}
