package apihelper

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/hodl-repos/pdf-invoice/pkg/standardisedError"
)

type Parser[T any] interface {
	Parse(string) (*T, error)
}

// define custom error to catch later in withErrorHandler
type GetUrlQueryValueError struct {
	standardisedError.StandardisedError
	Param string `json:"param"`
}

func (m *GetUrlQueryValueError) Error() string {
	return fmt.Sprintf("cannot load value '%v' from query", m.Param)
}

func (m *GetUrlQueryValueError) GetStandardisedError() *standardisedError.StandardisedError {
	return &m.StandardisedError
}

type UrlQueryTypes interface {
	string | int64 | float64 | uuid.UUID
}

func GetUrlQueryValue[T any](query url.Values, valueKey string, parser Parser[T]) (*T, error) {
	v := query.Get(valueKey)

	if len(v) == 0 {
		return nil, nil
	}

	res, err := parser.Parse(v)

	if err != nil {
		return nil, &GetUrlQueryValueError{
			Param:             valueKey,
			StandardisedError: generateStandardisedErrorQueryValue(),
		}
	}

	return res, nil
}

func GetAllUrlQueryValues[T any](query url.Values, valueKey string, parser Parser[T]) (*[]T, error) {
	v := query[valueKey]

	if len(v) == 0 {
		return nil, nil
	}

	resArray := make([]T, 0)

	for _, item := range v {
		res, err := parser.Parse(item)

		if err != nil {
			return nil, &GetUrlQueryValueError{
				Param:             valueKey,
				StandardisedError: generateStandardisedErrorQueryValue(),
			}
		}

		resArray = append(resArray, *res)
	}

	return &resArray, nil
}

func generateStandardisedErrorQueryValue() standardisedError.StandardisedError {
	return standardisedError.StandardisedError{
		Type:   "https://example.net/bad-request",
		Title:  "Could not parse the query-param",
		Status: http.StatusBadRequest,
		Detail: "Could not parse the query-param, is it the right type type?",
	}
}
