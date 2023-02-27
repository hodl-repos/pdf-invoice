package service

import (
	"context"
	"log"

	"contrib.go.opencensus.io/exporter/prometheus"
	"github.com/go-chi/chi"
	"github.com/hodl-repos/pdf-invoice/internal/serverenv"
	"github.com/hodl-repos/pdf-invoice/pkg/logging"
	"github.com/hodl-repos/pdf-invoice/pkg/middleware"
)

type Server struct {
	config *Config
	env    *serverenv.ServerEnv
}

func NewServer(ctx context.Context, cfg *Config, env *serverenv.ServerEnv) (*Server, error) {
	if env.Localize() == nil {
		panic("no localize configured")
	}

	return &Server{
		config: cfg,
		env:    env,
	}, nil
}

func (s *Server) Routes(ctx context.Context) *chi.Mux {
	logger := logging.FromContext(ctx).Named("service")

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger(logger))

	r.Use(middleware.ApplySharedCors())

	exporter, err := prometheus.NewExporter(prometheus.Options{})
	if err != nil {
		log.Fatal(err)
	}

	r.Handle("/metrics", exporter)
	r.Route("/v1", s.v1Router)

	return r
}
