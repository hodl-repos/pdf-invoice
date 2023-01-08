package apihelper

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/hodl-repos/pdf-invoice/pkg/jsonutil"
)

type CreateRessourceResponse struct {
	ID *uuid.UUID `json:"id"`
}

func WriteCreateRessourceResponse(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	w.Header().Set("location", fmt.Sprintf("%s/%s", r.URL.Host+r.URL.Path, id.String()))

	jsonutil.MarshalResponse(w, http.StatusCreated, &CreateRessourceResponse{
		ID: &id,
	})
}
