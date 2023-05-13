package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/services"
	"log"
	"net/http"
)

var actionService services.ActionService

func init() {
	actionService = services.NewActionService()
}

func UploadActions(c *gin.Context) {
	var actions *[]models.ExamAction
	err := c.ShouldBind(&actions)
	if err != nil {
		c.JSON(http.StatusBadRequest, Resp{Code: 1, Msg: "invalid params"})
		return
	}
	err = actionService.UploadActions(actions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{Code: 1, Msg: "failed to upload"})
		return
	}
	c.JSON(http.StatusOK, Resp{Code: 0, Msg: "success"})
}

func UploadAction(c *gin.Context) {
	action := &models.ExamAction{}
	err := c.ShouldBind(action)
	if err != nil {
		c.JSON(http.StatusBadRequest, Resp{Code: 1, Msg: "invalid params"})
		return
	}
	err = actionService.UploadAction(action)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{Code: 1, Msg: "failed to upload"})
		return
	}
	c.JSON(http.StatusOK, Resp{Code: 0, Msg: "success"})
}

func SelectActionById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, Resp{Code: 1, Msg: "invalid params"})
		return
	}
	action, err := actionService.QueryActionByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{Code: 1, Msg: "failed to query action"})
		return
	}
	if action == nil {
		c.JSON(http.StatusOK, Resp{Code: 1, Msg: "action not found"})
		return
	}
	c.JSON(http.StatusOK, Resp{Code: 0, Msg: "success", Data: action})
}

func SelectActionByExamAndStudentID(c *gin.Context) {
	student := c.Param("student_id")
	exam := c.Param("exam_id")
	if exam == "" || student == "" {
		c.JSON(http.StatusBadRequest, Resp{Code: 1, Msg: "invalid params"})
		return
	}
	actions, err := actionService.SelectActionByExamAndStudentID(exam, student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{Code: 1, Msg: "failed to query action"})
		return
	}
	c.JSON(http.StatusOK, Resp{Code: 0, Msg: "success", Data: actions})
}

func QueryAction(c *gin.Context) {
	query, _ := c.GetRawData()
	if query == nil {
		c.JSON(http.StatusBadRequest, Resp{Code: 1, Msg: "invalid params"})
		return
	}
	actions, err := actionService.QueryAction(string(query))
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{Code: 1, Msg: "failed to query action"})
		return
	}
	log.Println(actions)
	c.JSON(http.StatusOK, Resp{Code: 0, Msg: "success", Data: actions})
}

func ListActionInQuestion(c *gin.Context) {
	questionID := c.Param("question_id")
	studentID := c.Param("student_id")
	examID := c.Param("exam_id")
	if questionID == "" || studentID == "" || examID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}

	actions, err := actionService.ListActionInQuestion(examID, studentID, questionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{Code: 1, Msg: "failed to query action"})
		return
	}
	c.JSON(http.StatusOK, Resp{Code: 0, Msg: "success", Data: actions})
}
