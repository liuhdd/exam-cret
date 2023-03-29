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
func UploadAction(c *gin.Context) {
	action := &models.ExamAction{}
	err := c.ShouldBind(action)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "params illa"})
		return
	}
	err = actionService.UploadAction(action)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "action upload successfully"})
}

func SelectActionById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "miss param id"})
		return
	}
	action, err := actionService.QueryActionByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to select action"})
		return
	}
	if action == nil {
		c.JSON(http.StatusOK, gin.H{"message": "action not exists"})
		return
	}
	c.JSON(http.StatusOK, action)
}

func SelectActionByExamAndStudentID(c *gin.Context) {
	student := c.Query("exam_id")
	exam := c.Query("student_id")
	if exam == "" || student == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "miss query params"})
		return
	}
	actions, err := actionService.SelectActionByExamAndStudentID(exam, student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query actions"})
		return
	}
	c.JSON(http.StatusOK, actions)
}

func QueryAction(c *gin.Context) {
	query, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "argument miss"})
		return
	}
	actions, err := actionService.QueryAction(string(query))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query actions"})
		return
	}
	log.Println(actions)
	c.JSON(http.StatusOK, actions)
}
