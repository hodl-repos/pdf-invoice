package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hodl-repos/pdf-invoice/internal/service"
	"github.com/hodl-repos/pdf-invoice/internal/setup"
	"github.com/hodl-repos/pdf-invoice/pkg/environment"
	"github.com/hodl-repos/pdf-invoice/pkg/httpServer"
	"github.com/hodl-repos/pdf-invoice/pkg/logging"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	if _, err := os.Stat(".env"); err == nil {
		err := environment.ImportFromFile(".env")
		if err != nil {
			fmt.Println("environment error", err)
		}
	}

	logger := logging.NewLoggerFromEnv()
	ctx = logging.ContextWithLogger(ctx, logger)

	// recover from panics
	defer func() {
		done()
		if r := recover(); r != nil {
			logger.Fatalw("application panic", "panic", r)
		}
	}()

	err := realMain(ctx)
	done()

	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("successful shutdown")
	ctx.Done()
}

func realMain(ctx context.Context) error {
	logger := logging.FromContext(ctx)

	logger.Infow("server listening")

	var cfg service.Config
	env, err := setup.Setup(ctx, &cfg)
	if err != nil {
		return fmt.Errorf("setup.Setup: %w", err)
	}
	defer env.Close(ctx)

	server, err := service.NewServer(ctx, &cfg, env)
	if err != nil {
		return fmt.Errorf("service.NewServer: %w", err)
	}

	srv, err := httpServer.New(cfg.Port)
	if err != nil {
		return fmt.Errorf("httpServer.New: %w", err)
	}

	logger.Infow("server listening", "port", cfg.Port)

	return srv.ServeHTTPWithHandler(ctx, server.Routes(ctx))
}
