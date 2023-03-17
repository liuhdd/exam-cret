package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/liuhdd/exam-cret/application/models"
	"net/http"
)

func UploadAction(c *gin.Context) {
	action := &models.ExamAction{}
	err := c.ShouldBind(action)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "params illa"})
	}
	//todo: store to blockchain
	c.JSON(http.StatusOK, gin.H{"message": "action upload successfully"})
}

func SelectActionById(c *gin.Context) {
	s := "{\"action_id\":\"action1\",\"exam_id\":\"exam1\",\"student_id\":\"student1\",\"action_type\":1,\"action_time\":1678293569,\"question_id\":\"question1\",\"answer\":\"123123\"}"
	a := &models.ExamAction{}
	json.Unmarshal([]byte(s), a)

	c.JSON(http.StatusOK, a)
}

//通过学生id查询学生的所有操作
func SelectActionByStudentId(c *gin.Context) {
	s := "[{\"action_id\":\"action1\",\"exam_id\":\"exam1\",\"student_id\":\"student1\",\"action_type\":1,\"action_time\":1678293569,\"question_id\":\"question1\",\"answer\":\"123123\"},{\"action_id\":\"action2\",\"exam_id\":\"exam1\",\"student_id\":\"student1\",\"action_type\":1,\"action_time\":1678293569,\"question_id\":\"question2\",\"answer\":\"123123\"}]"
	var a []*models.ExamAction
	json.Unmarshal([]byte(s), &a)
	c.JSON(http.StatusOK, a)
}
