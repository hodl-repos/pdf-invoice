package apihelper

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hodl-repos/pdf-invoice/pkg/standardisedError"
)

// define custom error to catch later in withErrorHandler
type GetContextValueError struct {
	standardisedError.StandardisedError
	Param string `json:"param"`
}

func (m *GetContextValueError) Error() string {
	return fmt.Sprintf("cannot load context value from %v", m.Param)
}

func (m *GetContextValueError) GetStandardisedError() *standardisedError.StandardisedError {
	return &m.StandardisedError
}

func GetContextValue[T any](ctx context.Context, valueKey string) (*T, error) {
	v, ok := ctx.Value(valueKey).(T)

	if !ok {
		return nil, &GetContextValueError{
			Param:             valueKey,
			StandardisedError: generateStandardisedErrorMarshal(),
		}
	}

	return &v, nil
}

func generateStandardisedErrorMarshal() standardisedError.StandardisedError {
	return standardisedError.StandardisedError{
		Type:   "https://example.net/internal-error",
		Title:  "Could not load the route-param from the content[value]",
		Status: http.StatusInternalServerError,
		Detail: "Could not load the route-param from the content[value]",
	}
}
