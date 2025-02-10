package main

import (
	"backend/lib/container"
	"backend/lib/envconstant"
	"backend/lib/router"
	"context"
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func main() {
	c, err := container.New(map[string]string{
		envconstant.LogLevelEnvVar:      os.Getenv(envconstant.LogLevelEnvVar),
		envconstant.DatabaseURLEnvVar:   os.Getenv(envconstant.DatabaseURLEnvVar),
		envconstant.PortEnvVar:          os.Getenv(envconstant.PortEnvVar),
		envconstant.DBMaxIdleEnvVar:     os.Getenv(envconstant.DBMaxIdleEnvVar),
		envconstant.DBMaxOpenEnvVar:     os.Getenv(envconstant.DBMaxOpenEnvVar),
		envconstant.DBMaxLifeTimeEnvVar: os.Getenv(envconstant.DBMaxLifeTimeEnvVar),
		envconstant.DBMaxIdleTimeEnvVar: os.Getenv(envconstant.DBMaxIdleTimeEnvVar),
		envconstant.ZapConf:             os.Getenv(envconstant.ZapConf),
		envconstant.GormConf:            os.Getenv(envconstant.GormConf),
		envconstant.PprofEnable:         os.Getenv(envconstant.PprofEnable),
	})
	if err != nil {
		defer func() {
			fmt.Println("server initialization failed error: %w", err)
		}()
		panic("server initialization failed")
	}

	if c.PprofEnable() == "true" {
		c.Logger().Info("Enabling pprof")
		go func() {
			mux := http.NewServeMux()
			mux.HandleFunc("/debug/pprof/", pprof.Index)
			mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
			mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
			mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
			mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

			server := &http.Server{
				Addr:    ":6060",
				Handler: mux,
			}
			c.Logger().Info("starting pprof server on :6060")
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				c.Logger().Fatal("server failed", zap.Error(err))
			}
		}()

	}

	// Create router
	router := router.NewRouter(c.Logger(), EndPointConf(c))

	if err := Run(router, c); err != nil {
		fmt.Println("server error:", err)
		os.Exit(1)
	}
}

func Run(router *chi.Mux, c container.Container) error {

	// Configure server
	server := &http.Server{
		Addr:         ":" + c.Port(),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown
	go func() {
		c.Logger().Info("starting server", zap.String("port", c.Port()))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			c.Logger().Fatal("server failed", zap.Error(err))
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Shutdown server
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		c.Logger().Error("server shutdown failed", zap.Error(err))
		return fmt.Errorf("server shutdown failed: %w", err)
	}
	c.Logger().Info("server stopped")
	return nil
}
