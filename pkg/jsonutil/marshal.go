package jsonutil

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hodl-repos/pdf-invoice/pkg/standardisedError"
)

// old
func MarshalResponse(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("content-type", "application/json")

	if response == nil {
		w.WriteHeader(status)
		return
	}

	data, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	fmt.Fprintf(w, "%s", data)
}

// new
type MarshalResponseError struct {
	standardisedError.StandardisedError
}

func (v *MarshalResponseError) GetStandardisedError() *standardisedError.StandardisedError {
	return &v.StandardisedError
}

func (v *MarshalResponseError) Error() string {
	return "cannot serialize json for response body"
}

func MarshalResponseWithError(w http.ResponseWriter, status int, response interface{}) error {
	w.Header().Set("content-type", "application/json")

	if response == nil {
		w.WriteHeader(status)
		return nil
	}

	data, err := json.Marshal(response)
	if err != nil {
		return &MarshalResponseError{
			StandardisedError: generateStandardisedErrorMarshal(),
		}
	}

	w.WriteHeader(status)
	fmt.Fprintf(w, "%s", data)
	return nil
}

// sets the standardisedError fields for a new marshal-response-error
func generateStandardisedErrorMarshal() standardisedError.StandardisedError {
	return standardisedError.StandardisedError{
		Type:   "https://example.net/serialize-error",
		Title:  "The generated server response could not be serialized to a valid json object",
		Status: http.StatusInternalServerError,
		Detail: "The generated server response could not be serialized to a valid json object, get in contact with the application operator to get this issue fixed",
	}
}
