package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成request_id
		requestId := uuid.New().String()
		c.Set("request_id", requestId)
		c.Next()
	}
}