package apihelper

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/hodl-repos/pdf-invoice/pkg/jsonutil"
)

func MapUuidHeader(header string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return mapUuidHeaderFunc(next, header)
	}
}

func mapUuidHeaderFunc(next http.Handler, header string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(chi.URLParam(r, header))
		if err != nil {
			jsonutil.MarshalResponse(w, http.StatusBadRequest, err)
			return
		}

		ctx := context.WithValue(r.Context(), header, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func MapUuidQueries(w http.ResponseWriter, r *url.URL, query string, trg *[]uuid.UUID) error {
	for _, strId := range r.Query()[query] {
		id, err := uuid.Parse(strId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%s", err)
			return err
		}
		*trg = append(*trg, id)
	}
	return nil
}
