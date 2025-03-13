package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/syedomair/backend-microservices/lib/container"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// MockContainer is a mock implementation of the container.Container interface
type MockContainer struct {
	port             string
	logger           *zap.Logger
	db               *gorm.DB
	pprofEnable      string
	pointServicePool *container.ConnectionPool
}

func (m *MockContainer) Port() string {
	return m.port
}

func (m *MockContainer) Logger() *zap.Logger {
	return m.logger
}

func (c *MockContainer) Db() *gorm.DB {
	return c.db
}
func (c *MockContainer) PprofEnable() string {
	return c.pprofEnable
}
func (c *MockContainer) PointServicePool() container.ConnectionPoolInterface {
	return c.pointServicePool
}

func TestRun(t *testing.T) {
	// Create a mock logger
	logger, _ := zap.NewDevelopment()

	// Create a mock container
	mockContainer := &MockContainer{
		port:        "8080",
		logger:      logger,
		pprofEnable: "false",
	}

	// Create a mock router
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Create a channel to simulate the interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a goroutine
	go func() {
		err := Run(router, mockContainer)
		assert.NoError(t, err)
	}()

	// Wait for the server to start
	time.Sleep(100 * time.Millisecond)

	// Send a request to the server
	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Simulate an interrupt signal to stop the server
	quit <- syscall.SIGINT

	// Wait for the server to shut down
	time.Sleep(100 * time.Millisecond)
}
