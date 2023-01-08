// Package jsonutil provides common utilities for properly handling JSON
// payloads in HTTP requests.
package jsonutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hodl-repos/pdf-invoice/pkg/standardisedError"
)

const (
	// Max request size of 64KB. Prevents unnecessarily parsing JSON payloads that
	// are much larger than anticipated.
	maxBodyBytes = 64_000
)

// old
// Unmarshal provides a common implemetation of JSON unmarshalling with well
// defined error handling
func Unmarshal(w http.ResponseWriter, r *http.Request, data interface{}) (int, error) {
	if t := r.Header.Get("content-type"); len(t) < 16 || t[:16] != "application/json" {
		return http.StatusUnsupportedMediaType, fmt.Errorf("content-type is not application/json")
	}

	defer r.Body.Close()

	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	if err := d.Decode(data); err != nil {
		var syntaxErr *json.SyntaxError
		var unmarshalError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxErr):
			return http.StatusBadRequest, fmt.Errorf("malformed json at position %d", syntaxErr.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return http.StatusBadRequest, fmt.Errorf("malformed json")
		case errors.As(err, &unmarshalError):
			return http.StatusBadRequest, fmt.Errorf("invalid value %q at position %d", unmarshalError.Field, unmarshalError.Offset)
		case strings.HasPrefix(err.Error(), "json: unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return http.StatusBadRequest, fmt.Errorf("unknown field %s", fieldName)
		case errors.Is(err, io.EOF):
			return http.StatusBadRequest, fmt.Errorf("body must not be empty")
		case err.Error() == "http: request body too large":
			return http.StatusRequestEntityTooLarge, err
		default:
			return http.StatusInternalServerError, fmt.Errorf("failed to decode json: %w", err)
		}
	}
	if d.More() {
		return http.StatusBadRequest, fmt.Errorf("body must contain only one JSON object")
	}

	return http.StatusOK, nil
}

type UnmarshalError struct {
	standardisedError.StandardisedError
}

func (v *UnmarshalError) GetStandardisedError() *standardisedError.StandardisedError {
	return &v.StandardisedError
}

func (v *UnmarshalError) Error() string {
	return "cannot read from request body"
}

// Unmarshal provides a common implemetation of JSON unmarshalling with well
// defined error handling
func UnmarshalWithError(w http.ResponseWriter, r *http.Request, data interface{}) error {
	if t := r.Header.Get("content-type"); len(t) < 16 || t[:16] != "application/json" {
		return &UnmarshalError{
			StandardisedError: generateStandardisedErrorUnmarshal("content-type is not application/json",
				"wront content type cannot be deserialized",
				http.StatusUnsupportedMediaType),
		}
	}

	defer r.Body.Close()

	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	if err := d.Decode(data); err != nil {
		var syntaxErr *json.SyntaxError
		var unmarshalError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxErr):
			return &UnmarshalError{
				StandardisedError: generateStandardisedErrorUnmarshal("malformed json",
					fmt.Sprintf("malformed json at position %d", syntaxErr.Offset),
					http.StatusBadRequest),
			}
		case errors.Is(err, io.ErrUnexpectedEOF):
			return &UnmarshalError{
				StandardisedError: generateStandardisedErrorUnmarshal("malformed json",
					"malformed json",
					http.StatusBadRequest),
			}
		case errors.As(err, &unmarshalError):
			return &UnmarshalError{
				StandardisedError: generateStandardisedErrorUnmarshal("invalid value found in json",
					fmt.Sprintf("invalid value %q at position %d", unmarshalError.Field, unmarshalError.Offset),
					http.StatusBadRequest),
			}
		case strings.HasPrefix(err.Error(), "json: unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return &UnmarshalError{
				StandardisedError: generateStandardisedErrorUnmarshal("json field ist not known",
					fmt.Sprintf("unknown field %s", fieldName),
					http.StatusBadRequest),
			}
		case errors.Is(err, io.EOF):
			return &UnmarshalError{
				StandardisedError: generateStandardisedErrorUnmarshal("body must not be empty",
					"body must not be empty",
					http.StatusBadRequest),
			}
		case err.Error() == "http: request body too large":
			return &UnmarshalError{
				StandardisedError: generateStandardisedErrorUnmarshal("http: request body too large",
					"http: request body too large",
					http.StatusRequestEntityTooLarge),
			}
		default:
			return &UnmarshalError{
				StandardisedError: generateStandardisedErrorUnmarshal("failed to decode json",
					fmt.Sprintf("failed to decode json: %v", err.Error()),
					http.StatusInternalServerError),
			}
		}
	}

	if d.More() {
		return &UnmarshalError{
			StandardisedError: generateStandardisedErrorUnmarshal("body must contain only one JSON object",
				"body must contain only one JSON object",
				http.StatusBadRequest),
		}
	}

	return nil
}

// sets the standardisedError fields for a new marshal-response-error
func generateStandardisedErrorUnmarshal(title, detail string, status int) standardisedError.StandardisedError {
	return standardisedError.StandardisedError{
		Type:   "https://example.net/deserialize-error",
		Title:  title,
		Status: status,
		Detail: detail,
	}
}
