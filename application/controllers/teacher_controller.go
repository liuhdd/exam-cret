package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/services"
)

var teacherService services.TeacherService

func init() {
	teacherService = services.NewTeacherService()
}

func CreateTeacher(c *gin.Context) {
	var teacher models.Teacher
	err := c.ShouldBindJSON(&teacher)
	if err != nil {
		c.JSON(400, Resp{Code: 1, Msg: "params missed or illegal"})
		return
	}
	err = teacherService.CreateTeacher(&teacher)
	if err != nil {
		c.JSON(500, Resp{Code: 1, Msg: err.Error()})
		return
	}
	c.JSON(200, Resp{Code: 0, Msg: "success"})
}

func GetTeacherByID(c *gin.Context) {
	id := c.Param("id")
	teacher, err := teacherService.GetTeacherByID(id)
	if err != nil {
		c.JSON(500, Resp{Code: 1, Msg: err.Error()})
		return
	}
	c.JSON(200, Resp{Code: 0, Msg: "success", Data: teacher})
}

func GetAllTeachers(c *gin.Context) {
	teachers, err := teacherService.GetAllTeachers()
	if err != nil {
		c.JSON(500, Resp{Code: 1, Msg: err.Error()})
		return
	}
	c.JSON(200, Resp{Code: 0, Msg: "success", Data: teachers})
}

func UpdateTeacher(c *gin.Context) {
	var teacher models.Teacher
	err := c.ShouldBindJSON(&teacher)
	if err != nil {
		c.JSON(400, Resp{Code: 1, Msg: "params missed or illegal"})
		return
	}
	err = teacherService.UpdateTeacher(&teacher)
	if err != nil {
		c.JSON(500, Resp{Code: 1, Msg: err.Error()})
		return
	}
	c.JSON(200, Resp{Code: 0, Msg: "success"})
}

func DeleteTeacher(c *gin.Context) {
	id := c.Param("id")
	err := teacherService.DeleteTeacher(id)
	if err != nil {
		c.JSON(500, Resp{Code: 1, Msg: err.Error()})
		return
	}
	c.JSON(200, Resp{Code: 0, Msg: "success"})
}

func GetTeacherByName(c *gin.Context) {
	name := c.Query("name")
	teachers, err := teacherService.GetTeacherByName(name)
	if err != nil {
		c.JSON(500, Resp{Code: 1, Msg: err.Error()})
		return
	}
	c.JSON(200, Resp{Code: 0, Msg: "success", Data: teachers})
}
