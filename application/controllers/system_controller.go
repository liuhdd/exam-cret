package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhdd/exam-cret/application/services"
	"net/http"
	"os"
)

func ConfirmResume(c *gin.Context) {
	err := services.ResumeData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{Code: 1})
		return
	}
	c.JSON(http.StatusOK, Resp{Code: 0})
}

func BackupData(c *gin.Context) {
	err := services.BackupData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{Code: 1})
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"backup.sql")
	c.Header("Content-Transfer-Encoding", "utf-8")
	getwd, _ := os.Getwd()
	c.File(getwd + "/backup.sql")
}
func DataResume(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err,
		})
		return
	}
	err = c.SaveUploadedFile(file, "/tmp/exam.sql")
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{Code: 1, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Resp{Code: 0})
}
