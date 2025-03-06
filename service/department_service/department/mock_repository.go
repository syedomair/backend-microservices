package department

import (
	"github.com/syedomair/backend-microservices/models"
)

// MockRepository is a manual mock implementation of the Repository interface.
type MockRepository struct {
	GetAllDepartmentDBFunc func(limit, offset int, orderBy, sort string) ([]*models.Department, string, error)
}

func (m *MockRepository) GetAllDepartmentDB(limit, offset int, orderBy, sort string) ([]*models.Department, string, error) {
	return m.GetAllDepartmentDBFunc(limit, offset, orderBy, sort)
}
