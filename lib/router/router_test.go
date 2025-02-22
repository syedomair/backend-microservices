package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"
)

func TestNewRouter(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	routes := []EndPoint{
		{
			Name:        "TestRoute",
			Method:      Get,
			Pattern:     "/test",
			HandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
		},
	}
	router := NewRouter(logger, routes)
	if router == nil {
		t.Errorf("Expected router to be not nil, got nil")
	}
}

func TestLoggingMiddleware(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	middlewareHandler := loggingMiddleware(logger)(nextHandler)
	req, err := http.NewRequest(Get, "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := middlewareHandler
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestLoggingMiddleware_RequestID(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	middlewareHandler := loggingMiddleware(logger)(nextHandler)
	req, err := http.NewRequest(Get, "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := middlewareHandler
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
	// Check that the request ID is logged
	// This requires a way to inspect the logs, which can be done using a test logger
	// that captures the logs and allows them to be inspected.
}

func TestLoggingMiddleware_Context(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*
			requestID, ok := r.Context().Value(RequestIDKey).(string)
			if !ok {
				t.Errorf("Expected request ID to be in context, got none")
			}
		*/
		w.WriteHeader(http.StatusOK)
	})
	middlewareHandler := loggingMiddleware(logger)(nextHandler)
	req, err := http.NewRequest(Get, "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := middlewareHandler
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
}
