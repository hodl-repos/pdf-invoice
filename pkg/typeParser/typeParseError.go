package typeParser

import (
	"fmt"
	"net/http"

	"github.com/hodl-repos/pdf-invoice/pkg/standardisedError"
)

func generateStandardisedErrorMessage(typeValue string) error {
	return &standardisedError.StandardisedError{
		Type:   "https://example.net/bad-request",
		Title:  "Could not parse the given value",
		Status: http.StatusBadRequest,
		Detail: fmt.Sprintf("Could not parse the given value to %v", typeValue),
	}
}
