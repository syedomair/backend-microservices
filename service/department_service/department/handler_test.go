package department

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syedomair/backend-microservices/models"
	"go.uber.org/zap"
)

func TestGetAllDepartments_Success(t *testing.T) {
	// Arrange
	logger, _ := zap.NewDevelopment()
	mockRepo := &MockRepository{
		GetAllDepartmentDBFunc: func(limit, offset int, orderBy, sort string) ([]*models.Department, string, error) {
			return []*models.Department{{ID: "1", Name: "HR", Address: "123 Main St"}}, "1", nil
		},
	}

	controller := &Controller{
		Logger: logger,
		Repo:   mockRepo,
	}

	req, err := http.NewRequest("GET", "/departments", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	// Act
	controller.GetAllDepartments(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	// Additional assertions can be made on the response body if needed
}

func TestGetAllDepartments_InvalidQueryParams(t *testing.T) {
	// Arrange
	logger, _ := zap.NewDevelopment()
	mockRepo := &MockRepository{}

	controller := &Controller{
		Logger: logger,
		Repo:   mockRepo,
	}

	req, err := http.NewRequest("GET", "/departments?limit=invalid&offset=invalid", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	// Act
	controller.GetAllDepartments(rr, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetAllDepartments_RepositoryError(t *testing.T) {
	// Arrange
	logger, _ := zap.NewDevelopment()
	mockRepo := &MockRepository{
		GetAllDepartmentDBFunc: func(limit, offset int, orderBy, sort string) ([]*models.Department, string, error) {
			return nil, "", errors.New("repository error")
		},
	}

	controller := &Controller{
		Logger: logger,
		Repo:   mockRepo,
	}

	req, err := http.NewRequest("GET", "/departments", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	// Act
	controller.GetAllDepartments(rr, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
