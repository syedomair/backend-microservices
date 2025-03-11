package container

import (
	"context"
	"fmt"
	"strconv"

	pb "github.com/syedomair/backend-microservices/protos/point"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
)

const (
	LogLevel      = "LOG_LEVEL"
	DatabaseURL   = "DATABASE_URL"
	Port          = "PORT"
	DB            = "DB"
	DBMaxIdle     = "DB_MAX_IDLE"
	DBMaxOpen     = "DB_MAX_OPEN"
	DBMaxLifeTime = "DB_MAX_LIFE_TIME"
	DBMaxIdleTime = "DB_MAX_IDLE_TIME"
	ZapConf       = "ZAP_CONF"
	GormConf      = "GORM_CONF"
	PprofEnable   = "PPROF_ENABLE"
	Postgres      = "POSTGRES"
	Mysql         = "MYSQL"
)

// Container interface
type Container interface {
	Logger() *zap.Logger
	Db() *gorm.DB
	Port() string
	PprofEnable() string
	PointServiceClient() PointServiceClient
}

type container struct {
	logger               *zap.Logger
	db                   *gorm.DB
	port                 string
	pprofEnable          string
	environmentVariables map[string]string
	pointServiceClient   PointServiceClient
}

type PointServiceClient interface {
	GetUserPoints(ctx context.Context, in *pb.PointRequest, opts ...grpc.CallOption) (*pb.PointReply, error)
}

var _ Container = (*container)(nil)

func (c *container) Db() *gorm.DB {
	return c.db
}

func New(envVars map[string]string) (Container, error) {
	requiredKeys := []string{
		DatabaseURL,
		Port,
		ZapConf,
		GormConf,
		PprofEnable,
		DBMaxIdle,
		DBMaxOpen,
		DBMaxLifeTime,
		DBMaxIdleTime,
	}

	for _, key := range requiredKeys {
		if _, ok := envVars[key]; !ok {
			return nil, fmt.Errorf("missing mandatory envvar: %q", key)
		}
	}

	c := &container{environmentVariables: envVars}

	var err error
	c.db, err = c.dbSetup()
	if err != nil {
		return c, err
	}
	c.logger, err = c.loggerSetup()
	if err != nil {
		return c, err
	}
	c.port, err = c.portSetup()
	if err != nil {
		return c, err
	}
	c.pprofEnable, err = c.pprofEnableSetup()
	if err != nil {
		return c, err
	}

	conn, err := grpc.NewClient("point_service:8185", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("did not connect error: %v", err)
	}
	//defer conn.Close()
	c.pointServiceClient = pb.NewPointServerClient(conn)

	return c, nil
}

func (c *container) PointServiceClient() PointServiceClient {
	return c.pointServiceClient
}
func (c *container) dbSetup() (*gorm.DB, error) {
	if c.db != nil {
		return c.db, nil
	}

	dbMaxIdle, err := c.getIntEnvVar(DBMaxIdle)
	if err != nil {
		return nil, err
	}

	dbMaxOpen, err := c.getIntEnvVar(DBMaxOpen)
	if err != nil {
		return nil, err
	}

	dbMaxLifeTime, err := c.getIntEnvVar(DBMaxLifeTime)
	if err != nil {
		return nil, err
	}

	dbMaxIdleTime, err := c.getIntEnvVar(DBMaxIdleTime)
	if err != nil {
		return nil, err
	}

	strDatabaseURLEnvVar, err := c.getRequiredEnvVar(DatabaseURL)
	if err != nil {
		return nil, err
	}

	strGormConf, err := c.getRequiredEnvVar(GormConf)
	if err != nil {
		return nil, err
	}

	ca := NewDBConnectionAdapter(DB,
		strDatabaseURLEnvVar,
		dbMaxIdle,
		dbMaxOpen,
		dbMaxLifeTime,
		dbMaxIdleTime,
		strGormConf)
	db, err := ca.MakeConnection()
	if err != nil {
		return nil, err
	}
	c.db = db
	return db, nil
}

func (c *container) loggerSetup() (*zap.Logger, error) {
	if c.logger != nil {
		return c.logger, nil
	}

	zapConfig, err := c.getRequiredEnvVar(ZapConf)
	if err != nil {
		return nil, err
	}
	logger, err := NewLogger(zapConfig)
	if err != nil {
		return nil, err
	}
	c.logger = logger
	return logger, nil
}

func (c *container) Logger() *zap.Logger {
	return c.logger
}

func (c *container) portSetup() (string, error) {
	return c.getRequiredEnvVar(Port)
}

func (c *container) Port() string {
	return c.port
}

func (c *container) pprofEnableSetup() (string, error) {
	return c.getRequiredEnvVar(PprofEnable)
}

func (c *container) PprofEnable() string {
	return c.pprofEnable
}

func (c *container) getIntEnvVar(key string) (int, error) {
	strVal, err := c.getRequiredEnvVar(key)
	if err != nil {
		return 0, err
	}
	intVal, err := strconv.Atoi(strVal)
	if err != nil {
		return 0, fmt.Errorf("failed to convert %q to int: %w", strVal, err)
	}
	return intVal, nil
}

func (c *container) getRequiredEnvVar(key string) (string, error) {
	value, ok := c.environmentVariables[key]
	if !ok {
		return "", fmt.Errorf("missing mandatory envvar: %q", key)
	}
	return value, nil
}
