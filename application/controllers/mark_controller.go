package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/services"
	"github.com/liuhdd/exam-cret/application/services/dto"
)

var markService services.MarkService

func init() {
	markService = services.NewMarkService()
}

func GetQuestionScore(c *gin.Context) {
	question := &dto.Question{}
	err := c.ShouldBind(question)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "params illa"})
		return
	}
	score, err := markService.FindMarkByQuestionID(question.ExamID, question.StudentID, question.QuestionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query score"})
		return
	}
	c.JSON(http.StatusOK, score)
}

func UploadExamScore(c *gin.Context) {
	mark := &models.MarkAction{}
	c.ShouldBind(mark)
	err := markService.UploadMarkAction(mark)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload score"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "upload score successfully"})
}

