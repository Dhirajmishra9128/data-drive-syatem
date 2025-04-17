package controllers

import (
	"data-drive-system/config"
	"data-drive-system/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var files []models.File
	config.DB.Where("user_id = ?", userID).Find(&files)
	c.HTML(http.StatusOK, "dashboard.html", gin.H{"files": files})
}

func CreateFile(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	name := c.PostForm("name")
	ftype := c.PostForm("type")
	parentStr := c.PostForm("parent_id")
	var parentID *uint
	if parentStr != "" {
		pid, _ := strconv.Atoi(parentStr)
		id := uint(pid)
		parentID = &id
	}
	file := models.File{Name: name, Type: ftype, ParentID: parentID, UserID: userID}
	config.DB.Create(&file)
	c.Redirect(http.StatusSeeOther, "/dashboard")
}

func UpdateFile(c *gin.Context) {
	id := c.Param("id")
	newName := c.PostForm("name")
	var file models.File
	config.DB.First(&file, id)
	file.Name = newName
	config.DB.Save(&file)
	c.Redirect(http.StatusSeeOther, "/dashboard")
}

func DeleteFile(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.File{}, id)
	c.Redirect(http.StatusSeeOther, "/dashboard")
}
