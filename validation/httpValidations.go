package validation

import (
	"github.com/asaskevich/govalidator"
	"testes.com/packages/types"
)

func ValidateData(data *types.Category) (map[string]string, bool) {
	result, err := govalidator.ValidateStruct(data)
	if err != nil {
		errors := govalidator.ErrorsByField(err)
		validationErrors := make(map[string]string)
		for field, msg := range errors {
			validationErrors[field] = msg
		}
		return validationErrors, false
	}
	return nil, result
}
