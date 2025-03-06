package department

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syedomair/backend-microservices/models"
)

func TestGetAllDepartmentDB_Success(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{
		GetAllDepartmentDBFunc: func(limit, offset int, orderBy, sort string) ([]*models.Department, string, error) {
			return []*models.Department{{ID: "1", Name: "HR", Address: "123 Main St"}}, "1", nil
		},
	}

	// Act
	departments, count, err := mockRepo.GetAllDepartmentDB(10, 0, "name", "asc")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 1, len(departments))
	assert.Equal(t, "1", count)
	assert.Equal(t, "HR", departments[0].Name)
}

func TestGetAllDepartmentDB_Error(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{
		GetAllDepartmentDBFunc: func(limit, offset int, orderBy, sort string) ([]*models.Department, string, error) {
			return nil, "", errors.New("database error")
		},
	}

	// Act
	departments, count, err := mockRepo.GetAllDepartmentDB(10, 0, "name", "asc")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, departments)
	assert.Equal(t, "", count)
}
