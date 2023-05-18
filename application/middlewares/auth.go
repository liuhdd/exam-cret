package middlewares

import (
	"context"
	"fmt"
	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/redis/go-redis/v9"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

var whiteList = map[string]string{
	"/login":    "POST",
	"/registry": "POST",
	"/ping":     "GET",
}
var rdb *redis.Client

func init() {
	rdb = config.GetRedisClient()
}
func withinWhiteList(url *url.URL, method string) bool {
	queryUrl := strings.Split(fmt.Sprint(url), "?")[0]
	if _, ok := whiteList[queryUrl]; ok {
		if whiteList[queryUrl] == method {
			return true
		}
		return false
	}
	return false
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !withinWhiteList(c.Request.URL, c.Request.Method) {

			token := c.GetHeader("Authorization")
			if token == "" {
				c.JSON(http.StatusForbidden, "error: not logged in")
				c.Abort()
			}
			var user *models.User
			err := rdb.HGetAll(context.Background(), token).Scan(&user)
			if err != nil {
				panic(redis.ClientBlocked)
			}
			c.Set("user", user)
		}
		c.Next()
	}
}
