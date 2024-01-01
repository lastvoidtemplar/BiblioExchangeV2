package validation

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ValidationErrorsSlice struct {
	errors []error
}

func (sl *ValidationErrorsSlice) Add(err error) {
	sl.errors = append(sl.errors, err)
}
func (sl *ValidationErrorsSlice) Empty() bool {
	return len(sl.errors) == 0
}

type Validador interface {
	Valid() ValidationErrorsSlice
}

func HandleValidationErrors(c echo.Context, validErrs ValidationErrorsSlice) error {
	stringSlice := make([]string, 0)

	for _, err := range validErrs.errors {
		stringSlice = append(stringSlice, err.Error())
	}

	return c.JSON(http.StatusBadRequest, stringSlice)
}
