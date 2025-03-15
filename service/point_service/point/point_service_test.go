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
func (m *MockRepository) GetUserListPointsDB(userIDs []string) (map[string]int32, error) {
	args := m.Called(userIDs)
	return args.Get(0).(map[string]int32), args.Error(1)
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

func TestPointService_GetUserListPoints(t *testing.T) {
	t.Run("success - user list points retrieved", func(t *testing.T) {
		repo := new(MockRepository)
		logger := zaptest.NewLogger(t)
		service := NewPointService(repo, logger)

		userIDs := []string{"1", "2", "3"}
		mapUserPoints := make(map[string]int32)
		mapUserPoints["1"] = int32(12)
		mapUserPoints["2"] = int32(13)
		mapUserPoints["3"] = int32(14)
		repo.On("GetUserListPointsDB", userIDs).Return(mapUserPoints, nil)

		points, err := service.GetUserListPoints(userIDs)

		assert.NoError(t, err)
		assert.Equal(t, 3, len(points))
		assert.Equal(t, int32(13), points["2"])
		repo.AssertExpectations(t)
	})

	t.Run("error - repository failure", func(t *testing.T) {
		repo := new(MockRepository)
		logger := zaptest.NewLogger(t)
		service := NewPointService(repo, logger)

		userIDs := []string{"1", "2", "3"}
		mapUserPoints := make(map[string]int32)
		mapUserPoints["1"] = int32(12)
		mapUserPoints["2"] = int32(13)
		mapUserPoints["3"] = int32(14)
		repo.On("GetUserListPointsDB", userIDs).Return(mapUserPoints, errors.New("db error"))

		points, err := service.GetUserListPoints(userIDs)

		assert.Error(t, err)
		assert.Equal(t, int32(0), points["2"])
		assert.EqualError(t, err, "db error")
		repo.AssertExpectations(t)
	})
}
