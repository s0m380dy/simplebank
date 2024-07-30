package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/s0m380dy/simplebank/util"
)

var validCurrecny validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		//check cureency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}