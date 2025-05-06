package controller

import (
	"catalog/src"
	"catalog/src/service"

	"github.com/gin-gonic/gin"
)

type ConnectionCtrl struct {
	service *service.ConnectionService
}

func NewConnectionCtrl(service *service.ConnectionService) *ConnectionCtrl {
	return &ConnectionCtrl{
		service: service,
	}
}

func (ctrl *ConnectionCtrl) Register(router *gin.RouterGroup) error {
	connGroup := router.Group("/connections")

	connGroup.GET("/:id", ctrl.GetConnectionHandler())
	connGroup.POST("/create", ctrl.CreateConnectionHandler())
	return nil
}

func (ctrl *ConnectionCtrl) GetConnectionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if len(id) == 0 {
			c.JSON(400, gin.H{"error": "Connection ID is required"})
			return
		}

		conn, err := ctrl.service.GetConnectionById(c.Request.Context(), id)
		if err != nil {
			if err == src.ErrorSystemNotFound {
				c.JSON(404, gin.H{"error": "Connection not found"})
			} else {
				c.JSON(500, gin.H{"error": "Internal server error"})
			}
		}

		conn.Credentials = nil // Clear credentials for security reasons
		c.JSON(200, conn)
	}
}

func (ctrl *ConnectionCtrl) CreateConnectionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var conn service.Connection
		if err := c.ShouldBindJSON(&conn); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		_, err := ctrl.service.CreateConnection(c.Request.Context(), conn)
		if err != nil {
			if err == src.ErrorSystemNotFound {
				c.JSON(404, gin.H{"error": "Connection not found"})
			} else {
				c.JSON(500, gin.H{"error": "Internal server error"})
			}
		}

		c.Status(200)
	}
}
