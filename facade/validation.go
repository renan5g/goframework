package facade

import (
	"github.com/renan5g/goframework/contracts/validation"
)

func Validation() validation.Validation {
	return App().MakeValidation()
}
