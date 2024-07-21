package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/rohitvpatil0810/go-url-shortener-api/pkg/logger"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Logger.Info("Request", slog.String("method", r.Method), slog.String("path", r.URL.Path), slog.Duration("duration", time.Since(start)))
	})
}
