package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syedomair/backend-microservices/models"
)

func TestGetAllUserDB_Success(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{
		GetAllUserDBFunc: func(limit, offset int, orderBy, sort string) ([]*models.User, string, error) {
			return []*models.User{{ID: "1", Name: "John Doe", Age: 30, Salary: 50000.0}}, "1", nil
		},
	}

	// Act
	users, count, err := mockRepo.GetAllUserDB(10, 0, "name", "asc")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "1", count)
	assert.Equal(t, "John Doe", users[0].Name)
}

func TestGetAllUserDB_Error(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{
		GetAllUserDBFunc: func(limit, offset int, orderBy, sort string) ([]*models.User, string, error) {
			return nil, "", errors.New("database error")
		},
	}

	// Act
	users, count, err := mockRepo.GetAllUserDB(10, 0, "name", "asc")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, users)
	assert.Equal(t, "", count)
}

func TestGetUserHighAge_Success(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{
		GetUserHighAgeFunc: func() (int, error) {
			return 40, nil
		},
	}

	// Act
	highAge, err := mockRepo.GetUserHighAge()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 40, highAge)
}

func TestGetUserHighAge_Error(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{
		GetUserHighAgeFunc: func() (int, error) {
			return 0, errors.New("database error")
		},
	}

	// Act
	highAge, err := mockRepo.GetUserHighAge()

	// Assert
	assert.Error(t, err)
	assert.Equal(t, 0, highAge)
}

func TestGetUserLowAge_Success(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{
		GetUserLowAgeFunc: func() (int, error) {
			return 20, nil
		},
	}

	// Act
	lowAge, err := mockRepo.GetUserLowAge()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 20, lowAge)
}

func TestGetUserLowAge_Error(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{
		GetUserLowAgeFunc: func() (int, error) {
			return 0, errors.New("database error")
		},
	}

	// Act
	lowAge, err := mockRepo.GetUserLowAge()

	// Assert
	assert.Error(t, err)
	assert.Equal(t, 0, lowAge)
}

func TestGetUserAvgAge_Success(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{
		GetUserAvgAgeFunc: func() (float64, error) {
			return 30.5, nil
		},
	}

	// Act
	avgAge, err := mockRepo.GetUserAvgAge()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 30.5, avgAge)
}

func TestGetUserAvgAge_Error(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{
		GetUserAvgAgeFunc: func() (float64, error) {
			return 0, errors.New("database error")
		},
	}

	// Act
	avgAge, err := mockRepo.GetUserAvgAge()

	// Assert
	assert.Error(t, err)
	assert.Equal(t, 0.0, avgAge)
}

func TestGetUserLowSalary_Success(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{
		GetUserLowSalaryFunc: func() (float64, error) {
			return 30000.0, nil
		},
	}

	// Act
	lowSalary, err := mockRepo.GetUserLowSalary()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 30000.0, lowSalary)
}

func TestGetUserLowSalary_Error(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{
		GetUserLowSalaryFunc: func() (float64, error) {
			return 0, errors.New("database error")
		},
	}

	// Act
	lowSalary, err := mockRepo.GetUserLowSalary()

	// Assert
	assert.Error(t, err)
	assert.Equal(t, 0.0, lowSalary)
}

func TestGetUserHighSalary_Success(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{
		GetUserHighSalaryFunc: func() (float64, error) {
			return 100000.0, nil
		},
	}

	// Act
	highSalary, err := mockRepo.GetUserHighSalary()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 100000.0, highSalary)
}

func TestGetUserHighSalary_Error(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{
		GetUserHighSalaryFunc: func() (float64, error) {
			return 0, errors.New("database error")
		},
	}

	// Act
	highSalary, err := mockRepo.GetUserHighSalary()

	// Assert
	assert.Error(t, err)
	assert.Equal(t, 0.0, highSalary)
}

func TestGetUserAvgSalary_Success(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{
		GetUserAvgSalaryFunc: func() (float64, error) {
			return 65000.0, nil
		},
	}

	// Act
	avgSalary, err := mockRepo.GetUserAvgSalary()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 65000.0, avgSalary)
}

func TestGetUserAvgSalary_Error(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{
		GetUserAvgSalaryFunc: func() (float64, error) {
			return 0, errors.New("database error")
		},
	}

	// Act
	avgSalary, err := mockRepo.GetUserAvgSalary()

	// Assert
	assert.Error(t, err)
	assert.Equal(t, 0.0, avgSalary)
}
