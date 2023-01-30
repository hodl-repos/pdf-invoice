package service

import (
	"github.com/go-chi/chi"
	v1 "github.com/hodl-repos/pdf-invoice/internal/service/handler/v1"
	"github.com/hodl-repos/pdf-invoice/pkg/apihelper"
	errorhandling "github.com/hodl-repos/pdf-invoice/pkg/apihelper/errorHandling"
)

func (s *Server) v1Router(r chi.Router) {
	r.Get("/ping", apihelper.HandlePing())

	r.Get("/generate", errorhandling.WithError(v1.Handler()))
}
