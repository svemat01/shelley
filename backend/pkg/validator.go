package pkg

import (
	"github.com/go-playground/validator/v10"
	"log"
	"reflect"
	"strings"
)

type ErrorResponse struct {
	Field  string `json:"field"`
	Value  string `json:"value"`
	Reason string `json:"reason"`
}

var validate = validator.New()

func init() {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	log.Println("Validator initialized")
}

func ValidateStruct(data interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.Field()
			element.Reason = err.Tag()
			element.Value = err.Value().(string)
			errors = append(errors, &element)
		}
	}
	return errors
}
