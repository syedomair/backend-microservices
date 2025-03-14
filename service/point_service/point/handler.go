package point

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/syedomair/backend-microservices/lib/container"
	pb "github.com/syedomair/backend-microservices/proto/v1/point"
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
	server.grpcServer = grpc.NewServer()
	pb.RegisterPointServerServer(server.grpcServer, handler)

	return server, nil
}

// implement grpc server interface
//func (p *pointHandler) mustEmbedUnimplementedPointServerServer() {}

/*
package point

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/syedomair/backend-microservices/lib/container"
	pb "github.com/syedomair/backend-microservices/proto/v1/point"
)

var (
	netListen          = net.Listen
	jsonMarshal        = json.Marshal
	protojsonUnmarshal = protojson.Unmarshal
)

type Server interface {
	Serve() error
	GracefulStop()
}

type server struct {
	listener   net.Listener
	grpcServer *grpc.Server
	pb.UnimplementedPointServerServer
	container container.Container
}

func (s *server) Serve() error {
	return s.grpcServer.Serve(s.listener)
}

func (s *server) GracefulStop() {
	s.grpcServer.GracefulStop()
}

// GetUserPoints
func (s *server) GetUserPoints(_ context.Context, in *pb.PointRequest) (*pb.PointReply, error) {
	methodName := "GetUserPoints"
	s.container.Logger().Debug("method start", zap.String("method", methodName))
	start := time.Now()

	pointService := NewPointService(NewDBRepository(s.container.Db(), s.container.Logger()), s.container.Logger())
	userPoint, err := pointService.GetUserPoints(in.GetUserId())
	if err != nil {
		return nil, err
	}

	s.container.Logger().Debug("method end", zap.String("method", methodName), zap.Duration("duration", time.Since(start)))
	return &pb.PointReply{UserPoint: strconv.Itoa(userPoint)}, nil
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
	server.container = c
	server.grpcServer = grpc.NewServer()
	pb.RegisterPointServerServer(server.grpcServer, server)

	return server, nil
}
*/
