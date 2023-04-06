package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/services"
)

var markService services.MarkService

func init() {
	markService = services.NewMarkService()
}

func GetQuestionScore(c *gin.Context) {
	questionID := c.Query("question_id")
	studentID := c.Query("student_id")
	examID := c.Query("exam_id")
	if questionID == "" || studentID == "" || examID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "params illa"})
		return
	}
	score, err := markService.FindMarkByQuestionID(examID, studentID, questionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query score"})
		return
	}
	c.JSON(http.StatusOK, score)
}

func UploadExamScore(c *gin.Context) {
	mark := &models.MarkAction{}
	err := c.ShouldBind(mark)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "params illa"})
		return
	}
	err = markService.UploadMarkAction(mark)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload score"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "upload score successfully"})
}
