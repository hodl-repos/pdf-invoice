package qr

import (
	"github.com/go-chi/chi"
	"github.com/hodl-repos/pdf-invoice/internal/qr/handler"
	"github.com/hodl-repos/pdf-invoice/pkg/apihelper"
)

func (s *Server) qrV1Router(r chi.Router) {
	r.Get("/ping", apihelper.HandlePing())
	r.Get("/pdf", handler.HandlePdfGetV1())
}
