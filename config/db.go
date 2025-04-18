package config

import (
	"data-drive-system/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:MySecure@123@tcp(127.0.0.1:3306)/datadrive?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed", err)
	}
	DB = db
	DB.AutoMigrate(&models.User{}, &models.File{})
}
