package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Controller interface {
	Register(router *gin.RouterGroup) error
}

func RegisterApiRoutes(router *gin.RouterGroup, ctrls ...Controller) error {
	for _, ctrl := range ctrls {
		if err := ctrl.Register(router); err != nil {
			return errors.Wrap(err, "failed to register controller")
		}
	}
	return nil
}

func GetHealthCheckHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(200)
	}
}
