package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func CoolValidator(fl validator.FieldLevel) bool {
	return strings.Contains(fl.Field().String(), "cool")
}
