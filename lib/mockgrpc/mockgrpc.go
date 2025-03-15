package mockgrpc

import (
	"context"
	"errors"
	"net"
	"testing"

	pb "github.com/syedomair/backend-microservices/proto/v1/point"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type MockPointServiceClient struct {
	pb.UnimplementedPointServerServer
	pb.PointServerClient
}

func (m *MockPointServiceClient) GetUserPoints(ctx context.Context, in *pb.PointRequest) (*pb.PointReply, error) {
	return &pb.PointReply{UserPoint: "100"}, nil
}
func (m *MockPointServiceClient) GetUserListPoints(ctx context.Context, in *pb.UserListRequest) (*pb.UserListPointResponse, error) {
	mapUserPoints := make(map[string]int32)
	mapUserPoints["1"] = 10
	mapUserPoints["2"] = 20
	return &pb.UserListPointResponse{UserPoints: mapUserPoints}, nil
}

type MockConnectionPool struct {
	GetFunc func() (*grpc.ClientConn, error)
	PutFunc func(conn *grpc.ClientConn)
}

func (m *MockConnectionPool) Get() (*grpc.ClientConn, error) {
	if m.GetFunc != nil {
		return m.GetFunc()
	}
	return nil, errors.New("GetFunc not implemented")
}

func (m *MockConnectionPool) Put(conn *grpc.ClientConn) {
	if m.PutFunc != nil {
		m.PutFunc(conn)
	}
}
func (m *MockConnectionPool) Close() {}

func SetupGRPCServer(t *testing.T) (pb.PointServerClient, *grpc.ClientConn, *bufconn.Listener) {
	const bufSize = 1024 * 1024
	listener := bufconn.Listen(bufSize)
	srv := grpc.NewServer()
	pb.RegisterPointServerServer(srv, &MockPointServiceClient{}) // Register the mock service

	go func() {
		if err := srv.Serve(listener); err != nil {
			t.Errorf("Server exited with error: %v", err)
		}
	}()

	conn, err := grpc.NewClient(
		"passthrough:///bufnet",
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return listener.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	client := pb.NewPointServerClient(conn)

	return client, conn, listener
}
