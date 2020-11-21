package controllers

import (
	"net/http"

	"app/src/models"

	"github.com/gin-gonic/gin"
)

//UserController ...
type UserController struct{}

//Login ...
func (ctrl UserController) Login(c *gin.Context) {

	user := models.User{
		Email:    c.Query("email"),
		Password: c.Query("password"),
	}

	user, token, err := models.UserManager.Login(user)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "User signed in", "user": user, "token": token})
	} else {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid login details", "error": err.Error()})
	}

}
