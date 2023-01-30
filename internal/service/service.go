package service

import (
	"context"

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

	r.Route("/v1", s.v1Router)

	return r
}
