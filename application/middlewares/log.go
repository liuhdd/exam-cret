package middlewares

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	
        start := time.Now()

        path := c.Request.URL.Path

        method := c.Request.Method

        log.WithFields(log.Fields{
            "method": method,
            "path":   path,
        }).Info("request start")

        c.Next()

        end := time.Now()

        latency := end.Sub(start)

        log.WithFields(log.Fields{
            "method":   method,
            "path":     path,
            "status":   c.Writer.Status(),
            "duration": latency.Milliseconds(),
            "request_id": c.MustGet("request_id"),
        }).Info("request end")
	}
}