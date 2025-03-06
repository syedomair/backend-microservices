package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syedomair/backend-microservices/models"
)

type mockRepository struct {
	users      []*models.User
	count      string
	highAge    int
	lowAge     int
	highSalary float64
	lowSalary  float64
	avgAge     float64
	avgSalary  float64
}

func (m *mockRepository) GetAllUserDB(limit int, offset int, orderby string, sort string) ([]*models.User, string, error) {
	return m.users, m.count, nil
}
func (m *mockRepository) GetUserHighAge() (int, error) {
	return m.highAge, nil
}
func (m *mockRepository) GetUserLowAge() (int, error) {
	return m.lowAge, nil
}
func (m *mockRepository) GetUserAvgAge() (float64, error) {
	return m.avgAge, nil
}
func (m *mockRepository) GetUserLowSalary() (float64, error) {
	return m.lowSalary, nil
}
func (m *mockRepository) GetUserHighSalary() (float64, error) {
	return m.highSalary, nil
}
func (m *mockRepository) GetUserAvgSalary() (float64, error) {
	return m.avgSalary, nil
}

func TestGetAllUserDB_CustomMock(t *testing.T) {
	// Create a custom mock repository
	mockRepo := &mockRepository{
		users: []*models.User{
			{ID: "1", Name: "John Doe", Email: "john@example.com", DepartmentID: "dept1", Age: 30, Salary: 50000},
			{ID: "2", Name: "Jane Doe", Email: "jane@example.com", DepartmentID: "dept2", Age: 25, Salary: 60000},
		},
		count: "2",
	}

	// Call the function to test
	limit := 10
	offset := 0
	orderby := "name"
	sort := "asc"

	users, count, err := mockRepo.GetAllUserDB(limit, offset, orderby, sort)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 2, len(users))
	assert.Equal(t, "2", count)
}

func TestGetUserHighAge_CustomMock(t *testing.T) {
	// Create a custom mock repository
	mockRepo := &mockRepository{
		highAge: 50,
	}

	// Call the function to test
	highAge, err := mockRepo.GetUserHighAge()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 50, highAge)
}

func TestGetUserLowAge_CustomMock(t *testing.T) {
	// Create a custom mock repository
	mockRepo := &mockRepository{
		lowAge: 20,
	}

	// Call the function to test
	lowAge, err := mockRepo.GetUserLowAge()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 20, lowAge)
}

func TestGetUserAvgAge_CustomMock(t *testing.T) {
	// Create a custom mock repository
	mockRepo := &mockRepository{
		avgAge: 35.5,
	}

	// Call the function to test
	avgAge, err := mockRepo.GetUserAvgAge()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 35.5, avgAge)
}

func TestGetUserLowSalary_CustomMock(t *testing.T) {
	// Create a custom mock repository
	mockRepo := &mockRepository{
		lowSalary: 40000.0,
	}

	// Call the function to test
	lowSalary, err := mockRepo.GetUserLowSalary()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 40000.0, lowSalary)
}

func TestGetUserHighSalary_CustomMock(t *testing.T) {
	// Create a custom mock repository
	mockRepo := &mockRepository{
		highSalary: 100000.0,
	}

	// Call the function to test
	highSalary, err := mockRepo.GetUserHighSalary()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 100000.0, highSalary)
}

func TestGetUserAvgSalary_CustomMock(t *testing.T) {
	// Create a custom mock repository
	mockRepo := &mockRepository{
		avgSalary: 75000.0,
	}

	// Call the function to test
	avgSalary, err := mockRepo.GetUserAvgSalary()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 75000.0, avgSalary)
}
