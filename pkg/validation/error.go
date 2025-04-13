package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Errors []validator.FieldError

func As(errors validator.ValidationErrors) Errors {
	return Errors(errors)
}

func (ve Errors) Error() string {
	var resultError strings.Builder

	for _, fe := range ve {
		resultError.WriteString(fmt.Sprintf("%s should be a valid %s. ", fe.Field(), fe.Type()))
	}

	return strings.TrimSpace(resultError.String())
}
