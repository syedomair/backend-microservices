package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syedomair/backend-microservices/lib/mockcontainer"
	"github.com/syedomair/backend-microservices/models"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func TestGetAllUserStatistics_Success(t *testing.T) {
	// Setup mock repository
	mockRepo := &MockRepository{
		GetAllUserDBFunc: func(limit, offset int, orderBy, sort string) ([]*models.User, string, error) {
			return []*models.User{{ID: "1", Name: "John"}}, "100", nil
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

	pointServiceClient, conn, _ := mockcontainer.SetupGRPCServer(t) // Use the helper function
	defer conn.Close()

	mockConnectionPool := &mockcontainer.MockConnectionPool{
		GetFunc: func() (*grpc.ClientConn, error) {
			return conn, nil
		},
		PutFunc: func(conn *grpc.ClientConn) {
			// Do nothing in the mock
		},
	}

	// Initialize service with mock repository
	logger, _ := zap.NewProduction()
	userService := NewUserService(mockRepo, logger, pointServiceClient, mockConnectionPool)

	// Call the method under test
	result, err := userService.GetAllUserStatistics(10, 0, "id", "asc")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, &models.UserStatistics{
		UserList:       []*models.User{{ID: "1", Name: "John"}},
		Count:          "100",
		UserHighAge:    40,
		UserLowAge:     20,
		UserAvgAge:     30.5,
		UserLowSalary:  50000.0,
		UserHighSalary: 150000.0,
		UserAvgSalary:  100000.0,
	}, result)
}

func TestGetAllUserStatistics_ErrorInGetAllUserDB(t *testing.T) {

	mockRepo := &MockRepository{
		GetAllUserDBFunc: func(limit, offset int, orderBy, sort string) ([]*models.User, string, error) {
			return nil, "", errors.New("database error")
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

	pointServiceClient, conn, _ := mockcontainer.SetupGRPCServer(t) // Use the helper function
	defer conn.Close()

	mockConnectionPool := &mockcontainer.MockConnectionPool{
		GetFunc: func() (*grpc.ClientConn, error) {
			return conn, nil
		},
		PutFunc: func(conn *grpc.ClientConn) {
			// Do nothing in the mock
		},
	}
	// Initialize service with mock repository
	logger, _ := zap.NewProduction()
	userService := NewUserService(mockRepo, logger, pointServiceClient, mockConnectionPool)

	// Call the method under test
	result, err := userService.GetAllUserStatistics(10, 0, "id", "asc")

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
	assert.Nil(t, result)
}

func TestGetAllUserStatistics_ErrorInGetUserHighAge(t *testing.T) {

	mockRepo := &MockRepository{
		GetAllUserDBFunc: func(limit, offset int, orderBy, sort string) ([]*models.User, string, error) {
			return []*models.User{{ID: "1", Name: "John"}}, "100", nil
		},
		GetUserHighAgeFunc: func() (int, error) {
			return 0, errors.New("database error")
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

	pointServiceClient, conn, _ := mockcontainer.SetupGRPCServer(t) // Use the helper function
	defer conn.Close()

	mockConnectionPool := &mockcontainer.MockConnectionPool{
		GetFunc: func() (*grpc.ClientConn, error) {
			return conn, nil
		},
		PutFunc: func(conn *grpc.ClientConn) {
			// Do nothing in the mock
		},
	}
	// Initialize service with mock repository
	logger, _ := zap.NewProduction()
	userService := NewUserService(mockRepo, logger, pointServiceClient, mockConnectionPool)

	// Call the method under test
	result, err := userService.GetAllUserStatistics(10, 0, "id", "asc")

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
	assert.Nil(t, result)
}
