package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/services"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(as services.AuthService) *AuthController {
	return &AuthController{as}
}

// Register godoc
// @Summary Register
// @Description user register
// @Tags user
// @Accept  json
// @Produce  json
// @Param username body string true "username"
// @Param password body string true "password"
// @Success 200 {string} string	"register successfully"
// @Router /user/registry [post]
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

// Login godoc
// @Summary Login
// @Description user login
// @Tags user
// @Accept  json
// @Produce  json
// @Param username body string true "username"
// @Param password body string true "password"
// @Success 200 {string} string	"login successfully"
// @Router /user/login [post]
func (ac *AuthController) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "params missed or illegal"})
		return
	}
	err := ac.authService.Login(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to login. username or password is wrong"})
		return
	}
	session := sessions.Default(c)
	session.Set("uid", user.UserID)
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "login successfully"})
}

// Logout godoc
// @Summary Logout
// @Description user logout
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"logout successfully"
// @Router /user/logout [post]
func (ac *AuthController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("uid")
	err := session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to logout."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "logout successfully"})
}

// Ping godoc
// @Summary Ping
// @Description test connection
// @Tags ping
// @Accept  all
// @Produce  json
// @Success 200 {string} string	"pong"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
