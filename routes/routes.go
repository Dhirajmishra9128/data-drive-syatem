package routes

import (
	"data-drive-system/controllers"
	"data-drive-system/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) { c.HTML(200, "index.html", nil) })
	r.GET("/register", func(c *gin.Context) { c.HTML(200, "register.html", nil) })
	r.POST("/register", controllers.Register)
	r.GET("/login", func(c *gin.Context) { c.HTML(200, "login.html", nil) })
	r.POST("/login", controllers.Login)

	auth := r.Group("", middleware.AuthMiddleware())
	auth.GET("/dashboard", controllers.Dashboard)
	auth.POST("/drive/create", controllers.CreateFile)
	auth.POST("/drive/update/:id", controllers.UpdateFile)
	auth.POST("/drive/delete/:id", controllers.DeleteFile)
}
