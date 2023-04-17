package validation

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"testes.com/packages/types"
)

func ValidateData(data *types.Pessoa) (map[string]string, bool) {
	fmt.Println(data)
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
