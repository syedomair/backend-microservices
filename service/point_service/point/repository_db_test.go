package point

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetUserListPointsDB_Success(t *testing.T) {
	// Arrange
	mockRepo := &MockRepositoryDB{
		GetUserListPointsDBFunc: func(userIDs []string) (map[string]int32, error) {
			mapUserPoints := make(map[string]int32)
			mapUserPoints["1"] = int32(12)
			mapUserPoints["2"] = int32(13)
			mapUserPoints["3"] = int32(14)
			return mapUserPoints, nil
		},
	}

	userIDs := []string{"1", "2", "3"}
	// Act
	points, err := mockRepo.GetUserListPointsDB(userIDs)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 3, len(points))
	assert.Equal(t, int32(13), points["2"])
}

func TestGetUserPointDB_Success(t *testing.T) {
	// Arrange
	mockRepo := &MockRepositoryDB{
		GetUserPointDBFunc: func(userID string) (int, error) {
			return 1, nil
		},
	}
	// Act
	point, err := mockRepo.GetUserPointDB("10")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 1, point)
}

func TestGetUserPointDB_Error(t *testing.T) {
	// Arrange
	mockRepo := &MockRepositoryDB{
		GetUserPointDBFunc: func(userID string) (int, error) {
			return 0, errors.New("database error")
		},
	}

	// Act
	point, err := mockRepo.GetUserPointDB("10")

	// Assert
	assert.Error(t, err)
	assert.Equal(t, 0, point)
}
