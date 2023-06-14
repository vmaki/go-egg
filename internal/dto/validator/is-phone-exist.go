package validator

import "github.com/go-playground/validator/v10"

var IsPhoneExist validator.Func = func(fl validator.FieldLevel) bool {
	if fl.Field().String() != "15913395633" {
		return false
	}

	return true
}
