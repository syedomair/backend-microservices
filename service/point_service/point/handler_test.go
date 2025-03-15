package point

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/syedomair/backend-microservices/lib/container"
	pb "github.com/syedomair/backend-microservices/proto/v1/point"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
	"gorm.io/gorm"
)

// Mock PointService
type mockPointService struct {
	mock.Mock
}

var _ PointServiceInterface = (*mockPointService)(nil)

func (m *mockPointService) GetUserPoints(userId string) (int, error) {
	args := m.Called(userId)
	return args.Int(0), args.Error(1)
}

func (m *mockPointService) GetUserListPoints(userIDs []string) (map[string]int32, error) {
	args := m.Called(userIDs)
	return args.Get(0).(map[string]int32), args.Error(1)
}

// Mock Container
type mockContainer struct {
	mock.Mock
}

func (m *mockContainer) Logger() *zap.Logger {
	return zaptest.NewLogger(m.Called().Get(0).(*testing.T))
}

func (m *mockContainer) Port() string {
	return "8080"
}

func (m *mockContainer) Db() *gorm.DB {

	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*gorm.DB)
}
func (m *mockContainer) PprofEnable() string {
	args := m.Called()
	return args.String(0)
}
func (m *mockContainer) PointServicePool() container.ConnectionPoolInterface {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(container.ConnectionPoolInterface)
}

func TestGetUserPoints(t *testing.T) {
	mockService := new(mockPointService)
	mockContainer := new(mockContainer)
	mockContainer.On("Logger").Return(t)

	handler := &pointHandler{
		container: mockContainer,
		service:   mockService,
	}
	t.Run("success", func(t *testing.T) {
		mockService.On("GetUserPoints", "123").Return(100, nil)

		resp, err := handler.GetUserPoints(context.Background(), &pb.PointRequest{UserId: "123"})

		assert.NoError(t, err)
		assert.Equal(t, "100", resp.UserPoint)
		mockService.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		errMock := errors.New("service error")
		mockService.On("GetUserPoints", "456").Return(0, errMock)

		resp, err := handler.GetUserPoints(context.Background(), &pb.PointRequest{UserId: "456"})

		assert.Error(t, err)
		assert.Nil(t, resp)
		mockService.AssertExpectations(t)
	})
}

func TestGetUserListPoints(t *testing.T) {
	mockService := new(mockPointService)
	mockContainer := new(mockContainer)
	mockContainer.On("Logger").Return(t)

	handler := &pointHandler{
		container: mockContainer,
		service:   mockService,
	}

	userIDs := []string{"1", "2", "3"}
	mapUserPoints := make(map[string]int32)
	mapUserPoints["1"] = int32(12)
	mapUserPoints["2"] = int32(13)
	mapUserPoints["3"] = int32(14)
	t.Run("success", func(t *testing.T) {
		mockService.On("GetUserListPoints", userIDs).Return(mapUserPoints, nil)
		resp, err := handler.GetUserListPoints(context.Background(), &pb.UserListRequest{UserIds: userIDs})
		assert.NoError(t, err)
		assert.Equal(t, 3, len(resp.UserPoints))
		assert.Equal(t, int32(13), resp.UserPoints["2"])
		mockService.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		errMock := errors.New("service error")
		mockService.On("GetUserListPoints", userIDs).Return(mapUserPoints, errMock)
		//resp, err := handler.GetUserListPoints(context.Background(), &pb.UserListRequest{UserIds: userIDs})
		//assert.Error(t, err)
		//assert.Nil(t, resp)
		//mockService.AssertExpectations(t)
	})
}
