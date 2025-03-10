package main

import (
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"os"

	"github.com/syedomair/backend-microservices/lib/container"
	"go.uber.org/zap"

	server "github.com/syedomair/backend-microservices/service/point_service/pointserver"
)

func main() {
	c, err := container.New(map[string]string{
		container.LogLevel:      os.Getenv(container.LogLevel),
		container.DatabaseURL:   os.Getenv(container.DatabaseURL),
		container.Port:          os.Getenv(container.Port),
		container.DBMaxIdle:     os.Getenv(container.DBMaxIdle),
		container.DBMaxOpen:     os.Getenv(container.DBMaxOpen),
		container.DBMaxLifeTime: os.Getenv(container.DBMaxLifeTime),
		container.DBMaxIdleTime: os.Getenv(container.DBMaxIdleTime),
		container.ZapConf:       os.Getenv(container.ZapConf),
		container.GormConf:      os.Getenv(container.GormConf),
		container.PprofEnable:   os.Getenv(container.PprofEnable),
	})
	if err != nil {
		defer func() {
			log.Println("server initialization failed error: %w", err)
		}()
		panic("server initialization failed")
	}

	fmt.Println("DATABASE URL-------------------------------------------------", os.Getenv(container.DatabaseURL))
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	serverErrors := make(chan error, 1)

	server, err := server.NewServer(c)
	if err != nil {
		c.Logger().Fatal("failed to create server :", zap.Error(err))
	}

	go func() {
		c.Logger().Info("main: GRPC server listening on ", zap.String("port", c.Port()))
		serverErrors <- server.Serve()
	}()

	select {
	case err := <-serverErrors:
		c.Logger().Fatal("server error:", zap.Error(err))

	case sig := <-shutdown:
		c.Logger().Info("server shutdown:", zap.Any("signal", sig))
		server.GracefulStop()
	}

}
