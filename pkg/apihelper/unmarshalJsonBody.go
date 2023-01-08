package apihelper

import (
	"net/http"

	"github.com/hodl-repos/pdf-invoice/pkg/jsonutil"
	"github.com/hodl-repos/pdf-invoice/pkg/validation"
	"go.uber.org/zap"
)

// returns true when unmarshal succeeded, otherwise writes error to response and returns false
// expects a pointer to a dto
func UnmarshalJsonAndValidate(logger *zap.SugaredLogger, rw http.ResponseWriter, r *http.Request, data interface{}) bool {
	code, err := jsonutil.Unmarshal(rw, r, data)

	if err != nil {
		logger.Warnf("error unmarshaling API call, code: %v: %v", code, err.Error())

		if code == http.StatusInternalServerError {
			jsonutil.MarshalResponse(rw, http.StatusInternalServerError, nil)
			return false
		}

		type errorMsg struct {
			ErrorMessage string `json:"errorMessage"`
		}

		jsonutil.MarshalResponse(rw, code, &errorMsg{
			ErrorMessage: err.Error(),
		})

		return false
	}

	errorResponse := validation.ValidateStruct(data)
	if errorResponse != nil {
		jsonutil.MarshalResponse(rw, http.StatusBadRequest, errorResponse)
		return false
	}

	return true
}

// expects a pointer to a dto
func UnmarshalJsonAndValidateWithError(rw http.ResponseWriter, r *http.Request, data interface{}) error {
	err := jsonutil.UnmarshalWithError(rw, r, data)

	if err != nil {
		return err
	}

	err = validation.ValidateStruct(data)

	if err != nil {
		return err
	}

	return nil
}
