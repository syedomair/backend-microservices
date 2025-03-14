package point

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

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
