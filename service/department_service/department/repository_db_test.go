package department

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syedomair/backend-microservices/models"
)

type mockDepartmentRepository struct {
	departments []*models.Department
	count       string
}

func (m *mockDepartmentRepository) GetAllDepartmentDB(limit int, offset int, orderby string, sort string) ([]*models.Department, string, error) {
	return m.departments, m.count, nil
}

func TestGetAllDepartmentDB_CustomMock(t *testing.T) {
	// Create a custom mock repository
	mockRepo := &mockDepartmentRepository{
		departments: []*models.Department{
			{ID: "1", Name: "Test Depart 1"},
			{ID: "2", Name: "Test Depart 2"},
		},
		count: "2",
	}

	// Call the function to test
	limit := 10
	offset := 0
	orderby := "name"
	sort := "asc"

	departments, count, err := mockRepo.GetAllDepartmentDB(limit, offset, orderby, sort)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 2, len(departments))
	assert.Equal(t, "2", count)
}
