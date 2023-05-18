package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/controllers"
	"github.com/liuhdd/exam-cret/application/middlewares"
	swaggerfiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
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

	engine.Use(middlewares.Cors())
	engine.Use(middlewares.RequestIdMiddleware())
	engine.Use(middlewares.LogMiddleware())

	if config.GetProperty("environment") != "debug" {

		store := cookie.NewStore([]byte("secret"))
		engine.Use(sessions.Sessions("session", store))

		engine.Use(middlewares.AuthMiddleware())
	}

}

func SetupRoutes() {

	ac := controllers.NewAuthController()

	engine.GET("/ping", controllers.Ping)

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	user := engine.Group("user")
	{
		user.POST("/login", ac.Login)
		user.POST("/registry", ac.Register)
		user.POST("/logout", ac.Logout)
		user.GET("/list", ac.ListUsers)
		user.DELETE("/:username", ac.DeleteUser)
		user.Group("/")
	}

	system := engine.Group("system")
	{
		system.POST("/resume", controllers.DataResume)
		system.POST("/resume/confirm", controllers.ConfirmResume)
		system.GET("/backup", controllers.BackupData)
	}

	action := engine.Group("action")
	{
		action.POST("/upload", controllers.UploadAction)
		action.GET("/:id", controllers.SelectActionById)
		action.GET("/list/:exam_id/:student_id", controllers.SelectActionByExamAndStudentID)
		action.POST("/selector", controllers.QueryAction)
		action.GET("/question/:exam_id/:student_id/:question_id", controllers.ListActionInQuestion)
		action.POST("/uploads", controllers.UploadActions)
	}

	exam := engine.Group("exam")
	{
		exam.POST("/create", controllers.CreateExam)
		exam.GET("/show", controllers.ShowExamResult)
		exam.POST("/verify", controllers.VerifyExamResult)
		exam.DELETE("/:id", controllers.DeleteExam)
		exam.GET("/grade/:student_id", controllers.GetGradesByStudent)
		exam.GET("/list", controllers.ListExams)
		exam.POST("/query", controllers.QueryExams)
		exam.POST("/grades", controllers.QueryGrades)
	}

	score := engine.Group("score")
	{
		score.GET("/query", controllers.GetQuestionScore)
		score.POST("/upload", controllers.UploadExamScore)
	}

	teacher := engine.Group("teacher")
	{
		teacher.POST("/create", controllers.CreateTeacher)
		teacher.GET("/:id", controllers.GetTeacherByID)
		teacher.GET("/list", controllers.GetAllTeachers)
		teacher.POST("/update", controllers.UpdateTeacher)
		teacher.DELETE("/:id", controllers.DeleteTeacher)
		teacher.GET("/query", controllers.GetTeacherByName)
	}

	student := engine.Group("student")
	{
		student.POST("/create", controllers.CreateStudent)
		student.GET("/:id", controllers.GetStudentByID)
		student.GET("/list", controllers.GetAllStudents)
		student.POST("/update", controllers.UpdateStudent)
		student.DELETE("/:id", controllers.DeleteStudent)
		student.GET("/query", controllers.GetStudentByName)
	}

	question := engine.Group("question")
	{
		question.POST("/create", controllers.CreateQuestion)
		question.GET("/:id", controllers.GetQuestionByID)
	}
}
