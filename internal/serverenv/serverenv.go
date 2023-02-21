// Package serverenv defines common parameters for the server environment.
package serverenv

import (
	"context"

	"github.com/hodl-repos/pdf-invoice/pkg/localize"
	"github.com/hodl-repos/pdf-invoice/pkg/logging"
)

// ServerEnv represents latent environment configuration for servers in this
// application.
type ServerEnv struct {
	localizeService *localize.LocalizeService
}

// Option defines function type to modify a ServerEnv on creation.
type Option func(*ServerEnv) *ServerEnv

// New creates a new ServerEnv with the requested options.
func New(ctx context.Context, opts ...Option) *ServerEnv {
	env := &ServerEnv{}

	for _, o := range opts {
		env = o(env)
	}

	return env
}

func WithLocalization(service *localize.LocalizeService) Option {
	return func(env *ServerEnv) *ServerEnv {
		env.localizeService = service
		return env
	}
}

func (s *ServerEnv) Localize() *localize.LocalizeService {
	return s.localizeService
}

// Close shuts down the server env, closing database connections, etc.
func (s *ServerEnv) Close(ctx context.Context) error {
	logger := logging.FromContext(ctx)
	logger.Infow("serverenv.Close: running cleanup")

	if s == nil {
		return nil
	}

	return nil
}
