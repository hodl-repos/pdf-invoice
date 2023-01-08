package middleware

import (
	"net/http"

	"github.com/hodl-repos/pdf-invoice/pkg/logging"
	"go.uber.org/zap"
)

func Logger(originalLogger *zap.SugaredLogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			logger := originalLogger

			// Only override the logger if it's the default logger. This is only used
			// for testing and is intentionally a strict object equality check because
			// the default logger is a global default in the logger package.
			if existing := logging.FromContext(ctx); existing == logging.DefaultLogger() {
				logger = existing
			}

			// check if a request ID is set on the current context
			if id := RequestIDFromContext(ctx); id != "" {
				logger = logger.With(string(contextKeyRequestID), id)
			}

			// add logger to the current context
			ctx = logging.ContextWithLogger(ctx, logger)
			r = r.Clone(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
