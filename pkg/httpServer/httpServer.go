package httpServer

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/hodl-repos/pdf-invoice/pkg/logging"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/plugin/ochttp/propagation/tracecontext"
)

// Server provides a gracefolly-stoppable http server implementation. It is safe
// for concurrent use in goroutines
type Server struct {
	ip       string
	port     string
	listener net.Listener
}

// New creates a new server listening on the provided address that responds to
// the http.Handler. It starts the listener, but does not start the server. If
// an empty port is given, the server randomly chooses one.
func New(port string) (*Server, error) {
	// Create the net listener first, so the connection ready when we return. This
	// guarantees that it can accept requests.
	addr := fmt.Sprintf(":" + port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to create listener on %s: %w", addr, err)
	}

	return &Server{
		ip:       listener.Addr().(*net.TCPAddr).IP.String(),
		port:     strconv.Itoa(listener.Addr().(*net.TCPAddr).Port),
		listener: listener,
	}, nil
}

// ServeHTTP starts the server and blocks until the provided context is closed.
// When the provided context is closed, the server is gracefully stopped with a
// timeout of 5 seconds.
//
// Once a server has been stopped, it is NOT safe for reuse.
func (s *Server) ServeHTTP(ctx context.Context, srv *http.Server) error {
	logger := logging.FromContext(ctx)

	// Spawn a goroutine that listens for context closure. When the context is
	// closed, the server is stoped.
	errCh := make(chan error, 1)
	go func() {
		<-ctx.Done()

		logger.Debug("httpServer.ServeHTTP: context closed")
		shutdownCtx, done := context.WithTimeout(context.Background(), 5*time.Second)
		defer done()

		logger.Debug("httpServer.ServeHTTP: shutting down")
		if err := srv.Shutdown(shutdownCtx); err != nil {
			select {
			case errCh <- err:
			default:
			}
		}
	}()

	if err := srv.Serve(s.listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to serve: %w", err)
	}

	logger.Debug("httpServer.ServeHTTP: serving stopped")
	select {
	case err := <-errCh:
		return fmt.Errorf("failed to shutdown: %w", err)
	default:
		return nil
	}
}

func (s *Server) ServeHTTPWithHandler(ctx context.Context, handler http.Handler) error {
	return s.ServeHTTP(ctx, &http.Server{
		Handler: handler,
	})
}

// ServeHTTPHandler is a convenience wrapper around ServeHTTP. It creates an
// HTTP server using the provided handler, wrapped in OpenCensus for
// observability.
func (s *Server) ServeHTTPWithObservabilityHandler(ctx context.Context, handler http.Handler) error {
	return s.ServeHTTP(ctx, &http.Server{
		Handler: &ochttp.Handler{
			Handler:          handler,
			IsPublicEndpoint: true,
			Propagation:      &tracecontext.HTTPFormat{},
		},
	})
}
