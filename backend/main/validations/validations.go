package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func EntryTimeLength(fl validator.FieldLevel) bool {
	r := regexp.MustCompile("([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]")
	// -1 returns all matches, so we also need to make sure that is one
	matches := r.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	} else {
		return true
	}
}

func VarChar255Length(fl validator.FieldLevel) bool {
	if len(fl.Field().String()) > 255 {
		return false
	} else {
		return true
	}
}

// youtube custom validation
func UserRole(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case "TEACHER", "STUDENT", "ADMIN":
		return true
	}

	return false
}
