package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Temisaputra/warOnk/infrastructure/config"
	"github.com/Temisaputra/warOnk/pkg/auth"
	"github.com/Temisaputra/warOnk/pkg/helper"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type authMiddleware struct {
	cfg config.Config
}

func NewMiddleware(cfg config.Config) *authMiddleware {
	return &authMiddleware{cfg: cfg}
}

func (a *authMiddleware) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path, r.URL.Scheme)
		user, err := auth.ValidateCurrentUser(a.cfg, r)
		if err != nil {
			helper.WriteResponse(w, err, nil)
			return
		}

		r = auth.SetUserContext(r, user)
		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			rw := &statusResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			requestID := uuid.New().String()
			r = r.WithContext(context.WithValue(r.Context(), "request_id", requestID))

			next.ServeHTTP(rw, r)

			latency := time.Since(start)

			fields := []zap.Field{
				zap.String("request_id", requestID),
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Int("status_code", rw.statusCode),
				zap.Duration("latency", latency),
				zap.String("client_ip", r.RemoteAddr),
			}

			switch {
			case rw.statusCode >= 500:
				logger.Error("server error", fields...)
			case rw.statusCode >= 400:
				logger.Warn("client error", fields...)
			default:
				logger.Info("request handled", fields...)
			}
		})
	}
}

type statusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}
