package adapters

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

/*ValidErrToErrMap return error to map of interface*/
func ValidErrToErrMap(err error) {
	for _, err := range err.(validator.ValidationErrors) {
		fmt.Println(err.Namespace())
		fmt.Println(err.Field())
		fmt.Println(err.StructNamespace())
		fmt.Println(err.StructField())
		fmt.Println(err.Tag())
		fmt.Println(err.ActualTag())
		fmt.Println(err.Kind())
		fmt.Println(err.Type())
		fmt.Println(err.Value())
		fmt.Println(err.Param())
		fmt.Println()
	}

}
