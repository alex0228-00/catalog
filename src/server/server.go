package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"catalog/src"
	"catalog/src/controller"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Server struct {
	engine *gin.Engine
	server *http.Server
}

func NewServer(host, port string, ctrls ...controller.Controller) (*Server, error) {
	logger := src.DefaultLogger
	engine := gin.New()

	// Set up middleware
	logger.Info("Setting up server middleware")
	engine.Use(gin.Recovery())
	engine.Use(LongRequestLogger(src.DefaultLogger, time.Second*5))

	engine.GET("/health", controller.GetHealthCheckHandler())

	// Set up API versioning
	logger.Info("Registering API routes")
	v1 := engine.Group("/api/v1")
	if err := controller.RegisterApiRoutes(v1, ctrls...); err != nil {
		return nil, errors.Wrap(err, "failed to register API routes")
	}

	srv := &Server{
		engine: engine,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", host, port),
			Handler: engine,
		},
	}
	return srv, nil
}

func (srv *Server) ListenAndServe() chan error {
	ch := make(chan error, 1)
	go func() {
		src.DefaultLogger.Warn(fmt.Sprintf("Starting server on %s", srv.server.Addr))
		ch <- srv.server.ListenAndServe()
	}()
	return ch
}

func (srv *Server) Shutdown(ctx context.Context) error {
	if err := srv.server.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "failed to gracefully shutdown server")
	}
	return nil
}
