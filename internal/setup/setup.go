package setup

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"

	"github.com/hodl-repos/pdf-invoice/internal/serverenv"
	"github.com/hodl-repos/pdf-invoice/pkg/logging"
)

// Setup runs common initialization code for all servers. See SetupWith.
func Setup(ctx context.Context, config interface{}) (*serverenv.ServerEnv, error) {
	return SetupWith(ctx, config, envconfig.OsLookuper())
}

// SetupWith process the given configuration using envconfig. It is
// responsible for establishing a database connection, and accessing app
// configs. The provided interface must implement the various interfaces.
func SetupWith(ctx context.Context, config interface{}, l envconfig.Lookuper) (*serverenv.ServerEnv, error) {
	logger := logging.FromContext(ctx)

	// Build a list of options to pass to the server env.
	var serverEnvOpts []serverenv.Option

	// Process first round of environment variables.
	if err := envconfig.ProcessWith(ctx, config, l); err != nil {
		return nil, fmt.Errorf("error loading environment variables: %w", err)
	}
	logger.Infow("provided", "config", config)

	return serverenv.New(ctx, serverEnvOpts...), nil
}
