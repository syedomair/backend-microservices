package pointserver

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
	pb "github.com/syedomair/backend-microservices/service/user_service/pointserver/point"
)

/*
type Controller struct {
	Logger *zap.Logger
	Repo   Repository
}
*/

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
	//server.poetryDb = poetrydb.NewPoetryDb(config.PoetrydbBaseUrl, config.PoetrydbHttpTimeout)
	server.grpcServer = grpc.NewServer()
	//poetry.RegisterProtobufServiceServer(server.grpcServer, server)
	//reflection.Register(server.grpcServer)
	pb.RegisterPointServerServer(server.grpcServer, server)

	return server, nil
}

/*
func (s *server) RandomPoetries(ctx context.Context, in *poetry.RandomPoetriesRequest) (*poetry.PoetryList, error) {
	pbPoetryList := new(poetry.PoetryList)
	poetryList, err := s.poetryDb.Random(int(in.NumberOfPoetries))
	if err != nil {
		return pbPoetryList, errors.Wrap(err, "requesting random poetry")
	}
	json, err := jsonMarshal(poetryList)
	if err != nil {
		return pbPoetryList, errors.Wrap(err, "marshalling json")
	}
	err = protojsonUnmarshal(json, pbPoetryList)
	if err != nil {
		return pbPoetryList, errors.Wrap(err, "unmarshalling proto")
	}
	return pbPoetryList, nil
}

*/
