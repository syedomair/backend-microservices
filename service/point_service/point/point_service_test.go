package point

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zaptest"
)

// MockRepository is a mock implementation of the Repository interface
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetUserPointDB(userID string) (int, error) {
	args := m.Called(userID)
	return args.Int(0), args.Error(1)
}

func TestPointService_GetUserPoints(t *testing.T) {
	t.Run("success - points retrieved", func(t *testing.T) {
		repo := new(MockRepository)
		logger := zaptest.NewLogger(t)
		service := NewPointService(repo, logger)

		repo.On("GetUserPointDB", "user123").Return(100, nil)

		points, err := service.GetUserPoints("user123")

		assert.NoError(t, err)
		assert.Equal(t, 100, points)
		repo.AssertExpectations(t)
	})

	t.Run("error - repository failure", func(t *testing.T) {
		repo := new(MockRepository)
		logger := zaptest.NewLogger(t)
		service := NewPointService(repo, logger)

		repo.On("GetUserPointDB", "user123").Return(0, errors.New("db error"))

		points, err := service.GetUserPoints("user123")

		assert.Error(t, err)
		assert.Equal(t, 0, points)
		assert.EqualError(t, err, "db error")
		repo.AssertExpectations(t)
	})
}
