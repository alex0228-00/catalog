package server

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LongRequestLogger(logger *zap.Logger, timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		if latency := time.Since(start); latency > timeout {
			path := c.Request.URL.Path
			message := fmt.Sprintf("long request, url=[%s] %s, dur=%s", c.Request.Method, path, latency.String())
			logger.Warn(message)
		}
	}
}
