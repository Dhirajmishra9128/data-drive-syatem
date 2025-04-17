package controllers

import (
	"data-drive-system/config"
	"data-drive-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input models.User
	input.Username = c.PostForm("username")
	input.Password = c.PostForm("password")
	config.DB.Create(&input)
	c.Redirect(http.StatusSeeOther, "/login")
}

func Login(c *gin.Context) {
	var input models.User
	input.Username = c.PostForm("username")
	input.Password = c.PostForm("password")
	var user models.User
	config.DB.Where("username = ? AND password = ?", input.Username, input.Password).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	c.HTML(http.StatusOK, "token.html", gin.H{"token": user.ID})
}
