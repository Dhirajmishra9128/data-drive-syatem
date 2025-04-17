package main

import (
	"data-drive-system/config"
	"data-drive-system/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
