package api

import (
	"github.com/UnTea/WindingPacketsOnAPear/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupported(currency)
	}

	return false
}
