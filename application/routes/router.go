package routes

import (
	"github.com/liuhdd/exam-cret/application/middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/liuhdd/exam-cret/application/controllers"
	"github.com/liuhdd/exam-cret/application/repository"
	"github.com/liuhdd/exam-cret/application/services"
)



var engine *gin.Engine

func SetupRoutes() *gin.Engine {
	engine = gin.Default()
	store := cookie.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions("session", store))
	engine.Use(middlewares.AuthMiddleware())
	userRepository := repository.NewUserRepository()
	as := services.NewAuthService(userRepository)
	ac := controllers.NewAuthController(as)
	engine.GET("/ping", controllers.Ping)
	engine.POST("/login", ac.Login)
	engine.POST("/registry", ac.Register)
	engine.POST("/logout", ac.Logout)

	action := engine.Group("action")
	{
		action.POST("upload", controllers.UploadAction)
		action.POST("query/:id", controllers.SelectActionById)
		action.POST("queryByExamAndStudent", controllers.SelectActionByExamAndStudentID)
		action.POST("selector", controllers.QueryAction)
		action.POST("list", controllers.ListActionInQuestion)
	}
	
	return engine
}


