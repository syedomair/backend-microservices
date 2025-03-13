package user

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syedomair/backend-microservices/models"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func TestGetAllUsers_Success(t *testing.T) {
	// Arrange
	logger, _ := zap.NewDevelopment()
	mockRepo := &MockRepository{
		GetAllUserDBFunc: func(limit, offset int, orderBy, sort string) ([]*models.User, string, error) {
			return []*models.User{{ID: "1", Name: "John Doe", Age: 30, Salary: 50000.0}}, "1", nil
		},
		GetUserHighAgeFunc:    func() (int, error) { return 40, nil },
		GetUserLowAgeFunc:     func() (int, error) { return 20, nil },
		GetUserAvgAgeFunc:     func() (float64, error) { return 30.0, nil },
		GetUserHighSalaryFunc: func() (float64, error) { return 100000.0, nil },
		GetUserLowSalaryFunc:  func() (float64, error) { return 30000.0, nil },
		GetUserAvgSalaryFunc:  func() (float64, error) { return 65000.0, nil },
	}

	_, conn, _ := setupGRPCServer(t) // Use the helper function
	defer conn.Close()

	mockConnectionPool := &MockConnectionPool{
		GetFunc: func() (*grpc.ClientConn, error) {
			return conn, nil
		},
		PutFunc: func(conn *grpc.ClientConn) {
			// Do nothing in the mock
		},
	}

	controller := &Controller{
		Logger:                     logger,
		Repo:                       mockRepo,
		PointServiceConnectionPool: mockConnectionPool,
	}

	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	// Act
	controller.GetAllUsers(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	// Additional assertions can be made on the response body if needed
}

func TestGetAllUsers_InvalidQueryParams(t *testing.T) {
	// Arrange
	logger, _ := zap.NewDevelopment()
	mockRepo := &MockRepository{}

	controller := &Controller{
		Logger: logger,
		Repo:   mockRepo,
	}

	req, err := http.NewRequest("GET", "/users?limit=invalid&offset=invalid", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	// Act
	controller.GetAllUsers(rr, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetAllUsers_RepositoryError(t *testing.T) {
	// Arrange
	logger, _ := zap.NewDevelopment()
	mockRepo := &MockRepository{
		GetAllUserDBFunc: func(limit, offset int, orderBy, sort string) ([]*models.User, string, error) {
			return nil, "", errors.New("repository error")
		},
		GetUserHighAgeFunc: func() (int, error) {
			return 40, nil
		},
		GetUserLowAgeFunc: func() (int, error) {
			return 20, nil
		},
		GetUserAvgAgeFunc: func() (float64, error) {
			return 30.5, nil
		},
		GetUserLowSalaryFunc: func() (float64, error) {
			return 50000.0, nil
		},
		GetUserHighSalaryFunc: func() (float64, error) {
			return 150000.0, nil
		},
		GetUserAvgSalaryFunc: func() (float64, error) {
			return 100000.0, nil
		},
	}

	controller := &Controller{
		Logger: logger,
		Repo:   mockRepo,
	}

	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	// Act
	controller.GetAllUsers(rr, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetAllUsers_StatisticsError(t *testing.T) {
	// Arrange
	logger, _ := zap.NewDevelopment()
	mockRepo := &MockRepository{
		GetAllUserDBFunc: func(limit, offset int, orderBy, sort string) ([]*models.User, string, error) {
			return []*models.User{{ID: "1", Name: "John Doe", Age: 30, Salary: 50000.0}}, "1", nil
		},
		GetUserHighAgeFunc: func() (int, error) {
			return 0, errors.New("statistics error")
		},
		GetUserLowAgeFunc: func() (int, error) {
			return 20, nil
		},
		GetUserAvgAgeFunc: func() (float64, error) {
			return 30.5, nil
		},
		GetUserLowSalaryFunc: func() (float64, error) {
			return 50000.0, nil
		},
		GetUserHighSalaryFunc: func() (float64, error) {
			return 150000.0, nil
		},
		GetUserAvgSalaryFunc: func() (float64, error) {
			return 100000.0, nil
		},
	}

	_, conn, _ := setupGRPCServer(t) // Use the helper function
	defer conn.Close()

	mockConnectionPool := &MockConnectionPool{
		GetFunc: func() (*grpc.ClientConn, error) {
			return conn, nil
		},
		PutFunc: func(conn *grpc.ClientConn) {
			// Do nothing in the mock
		},
	}

	controller := &Controller{
		Logger:                     logger,
		Repo:                       mockRepo,
		PointServiceConnectionPool: mockConnectionPool,
	}

	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	// Act
	controller.GetAllUsers(rr, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
