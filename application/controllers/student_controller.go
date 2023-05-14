package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/services"
	log "github.com/sirupsen/logrus"
)

var studentService services.StudentService

func init() {
	studentService = services.NewStudentService()
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	err := c.ShouldBindJSON(&student)
	log.Debug(student)
	if err != nil {
		c.JSON(400, Resp{Code: 1, Msg: "params missed or illegal"})
		return
	}
	err = studentService.CreateStudent(&student)
	if err != nil {
		c.JSON(500, Resp{Code: 1, Msg: err.Error()})
		return
	}
	c.JSON(200, Resp{Code: 0, Msg: "success"})
}

func GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	student, err := studentService.GetStudentByID(id)
	if err != nil {
		c.JSON(500, Resp{Code: 1, Msg: err.Error()})
		return
	}
	c.JSON(200, Resp{Code: 0, Msg: "success", Data: student})
}

func GetAllStudents(c *gin.Context) {
	students, err := studentService.GetAllStudents()
	if err != nil {
		c.JSON(500, Resp{Code: 1, Msg: err.Error()})
		return
	}
	c.JSON(200, Resp{Code: 0, Msg: "success", Data: students})
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	err := c.ShouldBindJSON(&student)
	if err != nil {
		c.JSON(400, Resp{Code: 1, Msg: "params missed or illegal"})
		return
	}
	err = studentService.UpdateStudent(&student)
	if err != nil {
		c.JSON(500, Resp{Code: 1, Msg: err.Error()})
		return
	}
	c.JSON(200, Resp{Code: 0, Msg: "success"})
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	err := studentService.DeleteStudent(id)
	if err != nil {
		c.JSON(500, Resp{Code: 1, Msg: err.Error()})
		return
	}
	c.JSON(200, Resp{Code: 0, Msg: "success"})
}

func GetStudentByName(c *gin.Context) {
	name := c.Query("name")
	students, err := studentService.GetStudentByName(name)
	if err != nil {
		c.JSON(500, Resp{Code: 1, Msg: err.Error()})
		return
	}
	c.JSON(200, Resp{Code: 0, Msg: "success", Data: students})
}
