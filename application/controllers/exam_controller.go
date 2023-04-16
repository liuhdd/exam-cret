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
		c.JSON(400, gin.H{"error": "miss query params"})
		return
	}
	result, err := examService.FindExamResultByExamIDAndStudentID(examID, studentID)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to query exam result"})
		return
	}
	c.JSON(200, result)
}

func VerifyExamResult(c *gin.Context) {
	var result dto.ExamResult
	err := c.ShouldBind(&result)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	process, ok, err := examService.VerifyExamResults(&result)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to verify exam result"})
		return
	}
	c.JSON(200, gin.H{"process": process, "correct": ok})
}
