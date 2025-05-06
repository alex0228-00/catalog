package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"catalog/src"
	"catalog/src/controller"
	"catalog/src/datastore"
	"catalog/src/server"
	"catalog/src/service"

	"go.uber.org/zap"
)

func main() {
	logger := src.DefaultLogger
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)

	host := src.GetEnvOrPanic(src.EnvServerHost)
	port := src.GetEnvOrPanic(src.EnvServerPort)

	dsn := src.Dsn{
		Host:     src.GetEnvOrPanic(src.EnvDbHost),
		Port:     src.GetEnvOrPanic(src.EnvDbPort),
		Username: src.GetEnvOrPanic(src.EnvDbUser),
		Password: src.GetEnvOrPanic(src.EnvDbPassword),
		Database: src.GetEnvOrPanic(src.EnvDbSchema),
	}

	dbClient, err := datastore.GetDbClient(ctx, dsn)
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		os.Exit(1)
	}

	connSvc, err := service.NewConnectionService(nil, dbClient)
	if err != nil {
		logger.Error("Failed to create connection service", zap.Error(err))
		os.Exit(1)
	}

	srv, err := server.NewServer(host, port, controller.NewConnectionCtrl(connSvc))
	if err != nil {
		logger.Fatal("Failed to create server", zap.Error(err))
		os.Exit(1)
	}

	ch := srv.ListenAndServe()

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	cancel()

	select {
	case err := <-ch:
		logger.Error("Server error", zap.Error(err))
	case <-stopCh:
		// in case we cannot close the server
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		srv.Shutdown(ctx)
	}
}
