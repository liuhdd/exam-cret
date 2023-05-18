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
		c.JSON(http.StatusBadRequest, Resp{Code: 1, Msg: "invalid params"})
		return
	}
	score, err := markService.FindMarkByQuestionID(examID, studentID, questionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{Code: 1, Msg: "failed to query score"})
		return
	}
	c.JSON(http.StatusOK, Resp{Code: 0, Msg: "success", Data: score})
}

func UploadExamScores(c *gin.Context) {
	var marks []models.MarkAction
	err := c.ShouldBind(&marks)
	if err != nil {
		c.JSON(http.StatusBadRequest, Resp{Code: 1, Msg: "invalid params"})
		return
	}

	err = markService.SaveMarks(&marks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{Code: 1, Msg: "failed to upload score"})
		return
	}
	c.JSON(http.StatusOK, Resp{Code: 0, Msg: "success"})
}
func UploadExamScore(c *gin.Context) {
	mark := &models.MarkAction{}
	err := c.ShouldBind(mark)
	if err != nil {
		c.JSON(http.StatusBadRequest, Resp{Code: 1, Msg: "invalid params"})
		return
	}
	err = markService.UploadMarkAction(mark)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{Code: 1, Msg: "failed to upload score"})
		return
	}
	c.JSON(http.StatusOK, Resp{Code: 0, Msg: "success"})
}
