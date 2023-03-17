package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/services"
	"log"
	"net/http"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(as services.AuthService) *AuthController {
	return &AuthController{as}
}

func (ac *AuthController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ac.authService.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (ac *AuthController) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "params missed or illegal"})
		return
	}
	err := ac.authService.Login(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to login. username or password is wrong")})
		return
	}
	session := sessions.Default(c)
	session.Set("uid", user.UserID)
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "login successfully"})
}

func (ac *AuthController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("uid")
	err := session.Save()
	if err != nil {
		log.Printf("failed to save session")
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to logout.")})
		return
	}

}
