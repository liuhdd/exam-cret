package routes

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/liuhdd/exam-cret/application/controllers"
	"github.com/liuhdd/exam-cret/application/repository"
	"github.com/liuhdd/exam-cret/application/services"
	"net/http"
	"net/url"
	"strings"
)

var whiteList = map[string]string{
	"/login":    "POST",
	"/registry": "POST",
}

var engine *gin.Engine

func SetupRoutes() *gin.Engine {
	engine = gin.Default()
	store := cookie.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions("session", store))
	engine.Use(authMiddleware())
	userRepository := repository.NewUserRepository()
	as := services.NewAuthService(userRepository)
	ac := controllers.NewAuthController(as)
	engine.POST("/login", ac.Login)
	engine.POST("/registry", ac.Register)
	engine.POST("/logout", ac.Logout)

	action := engine.Group("action")
	{
		action.POST("upload", controllers.UploadAction)
		action.POST("query/:id", controllers.SelectActionById)
		action.POST("queryByExamAndStudent", controllers.SelectActionByExamAndStudentID)
		action.POST("selector", controllers.QueryAction)
	}
	return engine
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

func authMiddleware() gin.HandlerFunc {
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
