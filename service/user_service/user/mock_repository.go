package user

import (
	"github.com/syedomair/backend-microservices/models"
)

// MockRepository is a manual mock implementation of the Repository interface.
type MockRepository struct {
	GetAllUserDBFunc      func(limit, offset int, orderBy, sort string) ([]*models.User, string, error)
	GetUserHighAgeFunc    func() (int, error)
	GetUserLowAgeFunc     func() (int, error)
	GetUserAvgAgeFunc     func() (float64, error)
	GetUserLowSalaryFunc  func() (float64, error)
	GetUserHighSalaryFunc func() (float64, error)
	GetUserAvgSalaryFunc  func() (float64, error)
}

func (m *MockRepository) GetAllUserDB(limit, offset int, orderBy, sort string) ([]*models.User, string, error) {
	return m.GetAllUserDBFunc(limit, offset, orderBy, sort)
}

func (m *MockRepository) GetUserHighAge() (int, error) {
	return m.GetUserHighAgeFunc()
}

func (m *MockRepository) GetUserLowAge() (int, error) {
	return m.GetUserLowAgeFunc()
}

func (m *MockRepository) GetUserAvgAge() (float64, error) {
	return m.GetUserAvgAgeFunc()
}

func (m *MockRepository) GetUserLowSalary() (float64, error) {
	return m.GetUserLowSalaryFunc()
}

func (m *MockRepository) GetUserHighSalary() (float64, error) {
	return m.GetUserHighSalaryFunc()
}

func (m *MockRepository) GetUserAvgSalary() (float64, error) {
	return m.GetUserAvgSalaryFunc()
}
