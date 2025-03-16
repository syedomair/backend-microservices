package point

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	"github.com/syedomair/backend-microservices/lib/container"
	pb "github.com/syedomair/backend-microservices/proto/v1/point"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	netListen = net.Listen
)

type Server interface {
	Serve() error
	GracefulStop()
}

type PointHandler interface {
	GetUserPoints(ctx context.Context, in *pb.PointRequest) (*pb.PointReply, error)
	GetUserListPoints(ctx context.Context, in *pb.UserListRequest) (*pb.UserListPointResponse, error)
}

type server struct {
	listener   net.Listener
	grpcServer *grpc.Server
	handler    PointHandler
}

type pointHandler struct {
	container container.Container
	service   PointServiceInterface
	pb.UnimplementedPointServerServer
}

// GetUserPoints
func (p *pointHandler) GetUserPoints(_ context.Context, in *pb.PointRequest) (*pb.PointReply, error) {
	methodName := "GetUserPoints"
	p.container.Logger().Debug("method start", zap.String("method", methodName))
	start := time.Now()

	userPoint, err := p.service.GetUserPoints(in.GetUserId())
	if err != nil {
		return nil, err
	}

	p.container.Logger().Debug("method end", zap.String("method", methodName), zap.Duration("duration", time.Since(start)))
	return &pb.PointReply{UserPoint: strconv.Itoa(userPoint)}, nil
}

// GetUserListPoints
func (p *pointHandler) GetUserListPoints(_ context.Context, in *pb.UserListRequest) (*pb.UserListPointResponse, error) {
	methodName := "GetUserListPoints"
	p.container.Logger().Debug("method start", zap.String("method", methodName))
	start := time.Now()

	userPoint, err := p.service.GetUserListPoints(in.GetUserIds())
	if err != nil {
		return nil, err
	}

	p.container.Logger().Debug("method end", zap.String("method", methodName), zap.Duration("duration", time.Since(start)))
	return &pb.UserListPointResponse{UserPoints: userPoint}, nil
}

func (s *server) Serve() error {
	return s.grpcServer.Serve(s.listener)
}

func (s *server) GracefulStop() {
	s.grpcServer.GracefulStop()
}

// NewServer creates a new gRPC server.
func NewServer(c container.Container) (Server, error) {

	port, err := strconv.Atoi(c.Port())
	if err != nil {
		c.Logger().Fatal("invalid port value given:", zap.Error(err))
		return nil, err
	}

	server := new(server)
	listener, err := netListen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return server, errors.Wrap(err, "tcp listening")
	}
	server.listener = listener

	pointService := NewPointService(NewDBRepository(c.Db(), c.Logger()), c.Logger())
	handler := &pointHandler{
		container: c,
		service:   pointService,
	}

	server.handler = handler

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_prometheus.UnaryServerInterceptor,
		),
	)
	grpc_prometheus.Register(s)
	server.grpcServer = s

	//server.grpcServer = grpc.NewServer()
	pb.RegisterPointServerServer(server.grpcServer, handler)

	return server, nil
}
