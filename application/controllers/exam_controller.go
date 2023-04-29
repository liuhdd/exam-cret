package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhdd/exam-cret/application/services"
	"github.com/liuhdd/exam-cret/application/services/dto"
)

var examService services.ExamService

func init() {
	examService = services.NewExamService()
}

func ShowExamResult(c *gin.Context) {
	studentID := c.Query("student_id")
	examID := c.Query("exam_id")
	if examID == "" || studentID == "" {
		c.JSON(400, Resp{Code: 1, Msg: "invalid params"})
		return
	}
	result, err := examService.FindExamResultByExamIDAndStudentID(examID, studentID)
	if err != nil {
		c.JSON(500, Resp{Code: 1, Msg: "failed to query exam result"})
		return
	}
	c.JSON(200, Resp{Code: 0, Msg: "success", Data: result})
}

func VerifyExamResult(c *gin.Context) {
	var result dto.ExamResult
	err := c.ShouldBind(&result)
	if err != nil {
		c.JSON(400, Resp{Code: 1, Msg: "invalid params"})
		return
	}
	process, ok, err := examService.VerifyExamResults(&result)
	if err != nil {
		c.JSON(500, Resp{Code: 1, Msg: "failed to verify exam result"})
		return
	}
	c.JSON(200, Resp{Code: 0, Msg: "success", Data: map[string]interface{}{
		"process": process,
		"ok":      ok,
	}})
}
