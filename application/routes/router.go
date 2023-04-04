package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/controllers"
	"github.com/liuhdd/exam-cret/application/middlewares"
	"github.com/liuhdd/exam-cret/application/repository"
	"github.com/liuhdd/exam-cret/application/services"
)

var engine *gin.Engine
var inited bool

func InitEngine() *gin.Engine {
	if !inited {
		engine = gin.Default()
		SetUpMiddlewares()
		SetupRoutes()
	}
	return engine
}

func SetUpMiddlewares() {

	engine.Use(middlewares.RequestIdMiddleware())

	if config.GetProperty("environment") != "debug" {
		engine.Use(middlewares.LogMiddleware())
	}

	store := cookie.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions("session", store))

	engine.Use(middlewares.AuthMiddleware())

}

func SetupRoutes() {

	userRepository := repository.NewUserRepository()
	as := services.NewAuthService(userRepository)
	ac := controllers.NewAuthController(as)

	engine.GET("/ping", controllers.Ping)

	user := engine.Group("user")
	{
		user.POST("/login", ac.Login)
		user.POST("/registry", ac.Register)
		user.POST("/logout", ac.Logout)
	}

	action := engine.Group("action")
	{
		action.POST("upload", controllers.UploadAction)
		action.POST("query/:id", controllers.SelectActionById)
		action.POST("queryByExamAndStudent", controllers.SelectActionByExamAndStudentID)
		action.POST("selector", controllers.QueryAction)
		action.POST("list", controllers.ListActionInQuestion)
	}

	exam := engine.Group("exam")
	{
		exam.POST("result", controllers.ShowExamResult)

	}

}
