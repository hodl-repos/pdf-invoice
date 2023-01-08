package apihelper

import (
	"net/http"

	"github.com/hodl-repos/pdf-invoice/pkg/jsonutil"
	"github.com/hodl-repos/pdf-invoice/pkg/logging"
)

func HandlePing() func(w http.ResponseWriter, r *http.Request) {
	type pongResponse struct {
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		logger := logging.FromContext(ctx).Named("handlePingV1")

		logger.Debugw("got ping request")

		jsonutil.MarshalResponse(w, http.StatusOK, &pongResponse{
			Message: "pong",
		})
	}
}
