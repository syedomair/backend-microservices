package router

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	Get    = "GET"
	Patch  = "PATCH"
	Post   = "POST"
	Delete = "DELETE"
)

// EndPoint Public
type EndPoint struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type contextKey string

const (
	RequestIDKey contextKey = "requestID"
)

func NewRouter(logger *zap.Logger, routes []EndPoint) *chi.Mux {
	router := chi.NewRouter()

	// Common middleware
	router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second),
	)

	// Custom middleware
	router.Use(loggingMiddleware(logger))

	// Routes
	router.Route("/v1", func(r chi.Router) {
		for _, route := range routes {
			r.Method(route.Method, route.Pattern, route.HandlerFunc)
		}
	})

	// Health check
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return router
}
func loggingMiddleware(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			requestID := uuid.New().String()
			ctx := context.WithValue(r.Context(), RequestIDKey, requestID)

			logger.Info("request started",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.String("request_id", requestID),
			)

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			next.ServeHTTP(ww, r.WithContext(ctx))

			logger.Info("request completed",
				zap.String("request_id", requestID),
				zap.Int("status", ww.Status()),
				zap.Int("response_size", ww.BytesWritten()),
				zap.String("client_ip", r.RemoteAddr),
				zap.Duration("duration", time.Since(start)),
			)
		})
	}
}
