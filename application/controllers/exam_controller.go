package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/services"
	"github.com/liuhdd/exam-cret/application/services/dto"
)

var examService services.ExamService

func init() {
	examService = services.NewExamService()
}


func GetQuestionByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, Resp{Code: 1, Msg: "invalid params"})
		return
	}
	question, err := examService.FindQuestionByID(id)
	if err != nil {
		c.JSON(500, Resp{Code: 1, Msg: "failed to query question"})
		return
	}
	c.JSON(200, Resp{Code: 0, Msg: "success", Data: question})
}

func CreateQuestion(c *gin.Context) {
	var question models.Question
	err := c.ShouldBind(&question)
	if err != nil {
		c.JSON(400, Resp{Code: 1, Msg: "invalid params"})
		return
	}
	err = examService.SaveQuestion(&question)
	if err != nil {
		c.JSON(500, Resp{Code: 1, Msg: "failed to save question"})
		return
	}
	c.JSON(200, Resp{Code: 0, Msg: "success"})
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

func GetExamRecordByExamIDAndStudentID(c *gin.Context) {
	studentID := c.Query("student_id")
	examID := c.Query("exam_id")
	if examID == "" || studentID == "" {
		c.JSON(400, Resp{Code: 1, Msg: "invalid params"})
		return
	}
	result, err := examService.FindExamRecordByExamIDAndStudentID(examID, studentID)
	if err != nil {
		c.JSON(500, Resp{Code: 1, Msg: "failed to query exam records"})
		return
	}
	c.JSON(200, Resp{Code: 0, Msg: "success", Data: result})
}

func GetExamRecordsByStudentID(c *gin.Context) {
	studentID := c.Query("student_id")
	if studentID == "" {
		c.JSON(400, Resp{Code: 1, Msg: "invalid params"})
		return
	}
	result, err := examService.FindExamRecordsByStudentID(studentID)
	if err != nil {
		c.JSON(500, Resp{Code: 1, Msg: "failed to query exam records"})
		return
	}
	c.JSON(200, Resp{Code: 0, Msg: "success", Data: result})
}
