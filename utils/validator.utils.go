package utils

import (
	"github.com/go-playground/validator/v10"
)

// FormatValidationError mengubah error validator menjadi map[string]string yang rapi
func FormatValidationError(err error) (map[string]string, bool) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil, false
	}

	errorMessages := make(map[string]string)

	for _, e := range errs {
		field := e.Field()
		switch e.Tag() {
		case "required":
			errorMessages[field] = field + " wajib diisi."
		case "min":
			errorMessages[field] = field + " minimal " + e.Param() + " karakter."
		case "oneof":
			errorMessages[field] = field + " harus memilih antara 'Publish', 'Draft', atau 'Thrash'."
		default:
			errorMessages[field] = field + " tidak valid."
		}
	}

	return errorMessages, true
}
