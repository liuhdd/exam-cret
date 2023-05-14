package controllers

import (
	"net/http"
	"sync"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/services"
)

type AuthController struct {
	authService services.AuthService
}

var once sync.Once
var controller *AuthController

func NewAuthController() *AuthController {
	if controller == nil {
		once.Do(func() {
			controller = &AuthController{
				authService: services.NewAuthService(),
			}
		})
	}
	return controller
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
		c.JSON(http.StatusBadRequest, Resp{Code: 1, Msg: "params missed or illegal"})
		return
	}
	err := ac.authService.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{Code: 1, Msg: "failed to register"})
		return
	}
	c.JSON(http.StatusOK, Resp{Code: 0, Msg: "register successfully"})
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
		c.JSON(http.StatusBadRequest, Resp{Code: 1, Msg: "params missed or illegal"})
		return
	}
	err := ac.authService.Login(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, Resp{Code: 1, Msg: "failed to login"})
		return
	}
	session := sessions.Default(c)
	session.Set("uid", user.UserID)
	session.Save()
	c.JSON(http.StatusOK, Resp{Code: 0, Msg: "success"})
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
		c.JSON(http.StatusInternalServerError, Resp{Code: 1, Msg: "failed to logout"})
		return
	}
	c.JSON(http.StatusOK, Resp{Code: 0, Msg: "success"})
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
	c.JSON(http.StatusOK, "pong")
}
