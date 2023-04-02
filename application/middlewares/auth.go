package middlewares

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


var whiteList = map[string]string{
	"/login":    "POST",
	"/registry": "POST",
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
			session := sessions.Default(c)
			uid := session.Get("uid")
			if uid == nil {
				c.JSON(http.StatusForbidden, "error: not logged in")
				c.Abort()
			}
			c.Set("uid", uid)
		}
		c.Next()
	}
}