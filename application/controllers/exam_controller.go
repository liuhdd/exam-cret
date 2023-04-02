package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhdd/exam-cret/application/services"
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