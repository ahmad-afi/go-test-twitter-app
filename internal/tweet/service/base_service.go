package service

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/alwisisva/twitter-app/pkg/helper"
	"github.com/alwisisva/twitter-app/pkg/validator"
	v10 "github.com/go-playground/validator/v10"
)

func usecaseValidation[T any](params T) *helper.ErrorStruct {
	var validate = v10.New()

	if err := validate.Struct(params); err != nil {
		errs := err.(v10.ValidationErrors)
		var newErr validator.ValidationError

		for _, val := range errs {
			// fmt.Printf("error :  %#v \n\n ", val)
			// fmt.Printf("error :  %#v \n\n ", val.StructNamespace() )

			getField, _ := reflect.TypeOf(params).FieldByName(val.Field())
			jsonTag := getField.Tag.Get("json")

			var message string
			switch val.Tag() {
			case "required":
				message = fmt.Sprintf("%s cannot be null", jsonTag)
			default:
				message = fmt.Sprintf("validation error for '%s', Tag: %s", jsonTag, val.Tag())
			}

			newErr.Message = append(newErr.Message, message)
		}

		return &helper.ErrorStruct{
			Code: http.StatusBadRequest,
			Err:  newErr,
		}
	}

	return nil
}
