package validation

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/hodl-repos/pdf-invoice/pkg/standardisedError"
)

//#region ERROR struct

type InvalidParam struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

func (i *InvalidParam) Error() string {
	return fmt.Sprintf("%s %s", i.Name, i.Reason)
}

type ValidationError struct {
	standardisedError.StandardisedError
	InvalidParams []InvalidParam `json:"invalid-params"`
}

func (v *ValidationError) GetStandardisedError() *standardisedError.StandardisedError {
	return &v.StandardisedError
}

func (v *ValidationError) Error() string {
	var strs []string

	strs = append(strs, "validation failed")

	for _, invalidParam := range v.InvalidParams {
		strs = append(strs, invalidParam.Error())
	}

	return strings.Join(strs, "\n")
}

//#endregion ERROR struct

// validates the struct with go-playground validator, returns error when not valid
func ValidateStruct(data interface{}) error {
	err := validator.New().Struct(data)

	if err != nil {
		valErr := ValidationError{
			StandardisedError: generateStandardisedError(),
		}
		valErr.InvalidParams = make([]InvalidParam, 0)

		//combine all errors
		for _, err := range err.(validator.ValidationErrors) {
			invalidParam := InvalidParam{
				Name: strings.ToLower(err.Field()),
			}

			switch err.Tag() {
			case "required":
				invalidParam.Reason = "is required"
			default:
				invalidParam.Reason = fmt.Sprintf("failed on %s validation", err.Tag())
			}

			valErr.InvalidParams = append(valErr.InvalidParams, invalidParam)
		}

		return &valErr
	}

	return nil
}

// sets the standardisedError fields for a new validation-error
func generateStandardisedError() standardisedError.StandardisedError {
	return standardisedError.StandardisedError{
		Type:   "https://example.net/validation-error",
		Title:  "Your request body didn't validate",
		Status: http.StatusBadRequest,
		Detail: "Some properties of the given data was not valid for the endpoint, further information can be found in the invalid-params field",
	}
}

//no longer needed, errors already occur when parsing a uuid
// func ValidateUUID(s string, fieldname string) error {
// 	_, err := uuid.Parse(s)

// 	if err != nil {
// 		valErr := ValidationError{}
// 		valErr[fieldname] = []string{"invalid uuid"}

// 		return &ValidationErrorResponse{
// 			ValidationError: valErr,
// 		}
// 	}

// 	return nil
// }
