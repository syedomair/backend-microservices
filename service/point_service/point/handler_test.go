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

/*
import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/syedomair/backend-microservices/lib/container"
	pb "github.com/syedomair/backend-microservices/proto/v1/point"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"gorm.io/gorm"
)

// Mock Container
type MockContainer struct {
	mock.Mock
}

func (m *MockContainer) Db() *gorm.DB { // Corrected return type

	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*gorm.DB)
}

func (m *MockContainer) Logger() *zap.Logger {
	args := m.Called()
	return args.Get(0).(*zap.Logger)
}

func (m *MockContainer) Port() string {
	args := m.Called()
	return args.String(0)
}
func (m *MockContainer) PprofEnable() string {
	args := m.Called()
	return args.String(0)
}
func (m *MockContainer) PointServicePool() container.ConnectionPoolInterface {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(container.ConnectionPoolInterface)
}

// Mock Repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetUserPointDB(userID string) (int, error) {
	args := m.Called(userID)
	return args.Int(0), args.Error(1)
}

// Mock PointService
type MockPointService struct {
	PointService
	mock.Mock
}

func (m *MockPointService) GetUserPoints(userID string) (int, error) {
	args := m.Called(userID)
	return args.Int(0), args.Error(1)
}

func TestPointHandler_GetUserPoints(t *testing.T) {
	mockService := new(MockPointService)
	mockContainer := new(MockContainer)
	logger, _ := zap.NewDevelopment()
	mockContainer.On("Logger").Return(logger)

	handler := &pointHandler{
		container: mockContainer,
		service:   mockService, // Pass a pointer!
	}

	req := &pb.PointRequest{UserId: "test-user"}
	mockService.On("GetUserPoints", "test-user").Return(100, nil)

	reply, err := handler.GetUserPoints(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "100", reply.GetUserPoint())
	mockService.AssertExpectations(t)
}
*/
/*
func TestNewServer_Success(t *testing.T) {
	mockContainer := new(MockContainer)
	logger, _ := zap.NewDevelopment()
	mockContainer.On("Logger").Return(logger)
	mockContainer.On("Port").Return("50051")
	server, err := NewServer(mockContainer)
	assert.NoError(t, err)
	assert.NotNil(t, server)
	server.GracefulStop()
}
*/

/*
func TestNewServer_InvalidPort(t *testing.T) {
	mockContainer := new(MockContainer)
	logger, _ := zap.NewDevelopment()
	mockContainer.On("Logger").Return(logger)
	mockContainer.On("Port").Return("invalid_port")

	server, err := NewServer(mockContainer)

	assert.Error(t, err)
	assert.Nil(t, server)
}
*/
/*
func TestServer_ServeAndGracefulStop(t *testing.T) {
	lis := bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()
	server := &server{
		listener:   lis,
		grpcServer: s,
	}

	go func() {
		if err := server.Serve(); err != nil {
			t.Errorf("Server.Serve() error = %v", err)
		}
	}()

	time.Sleep(100 * time.Millisecond) // Give the server time to start

	server.GracefulStop()
}
*/
/*
func TestNewServer_ListenError(t *testing.T) {
	mockContainer := new(MockContainer)
	logger, _ := zap.NewDevelopment()
	mockContainer.On("Logger").Return(logger)
	mockContainer.On("Port").Return("50051")

	netListen = func(network, address string) (net.Listener, error) {
		return nil, errors.New("listen error")
	}
	defer func() {
		netListen = net.Listen // Restore original netListen
	}()

	server, err := NewServer(mockContainer)

	assert.Error(t, err)
	assert.Nil(t, server)
}
*/
