package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func LoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		total := time.Since(start)
		slog.Info("Incoming request", "method", r.Method, "path", r.URL.Path, "duration", total.String())
	})
}
