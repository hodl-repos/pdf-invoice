package errorhandling

import (
	"net/http"

	"github.com/hodl-repos/pdf-invoice/pkg/apihelper"
	"github.com/hodl-repos/pdf-invoice/pkg/jsonutil"
	"github.com/hodl-repos/pdf-invoice/pkg/logging"
	"github.com/hodl-repos/pdf-invoice/pkg/standardisedError"
)

// handles internal server errors for a http handler, should be used in every route
func WithError(f apihelper.HandlerFuncWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := logging.FromContext(ctx).Named("with-error")

		err := f(w, r)

		if err != nil {
			sdtError, ok := err.(GetStandardisedErrorInterface)
			if ok {
				errorModel := sdtError.GetStandardisedError()

				errorModel.Instance = r.URL.Host + r.URL.Path

				logger.Error(err.Error())
				//check if response can be serialized, otherwise return internal error without data
				if jsonutil.MarshalResponseWithError(w, errorModel.Status, err) == nil {
					return
				}
			}

			logger.Errorf("unexpected error from http handler: %v", err.Error())
			jsonutil.MarshalResponse(w, http.StatusInternalServerError, nil)
		}
	}
}

func wrapGormNotFoundError() error {
	return &standardisedError.StandardisedError{
		Type:   "https://example.net/not-found",
		Title:  "The requested ressource was not found",
		Status: http.StatusNotFound,
		Detail: "The requested ressource was not found in the database, have you lost it?",
	}
}
