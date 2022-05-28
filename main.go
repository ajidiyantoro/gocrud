package main

import (
	"log"

	"gocrud/handler"
	"gocrud/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Database Connection String
	dsn := "root:Password12345@tcp(127.0.0.1:3388)/gocrud?charset=utf8mb4&parseTime=True&loc=Local"

	// Connect to Database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.GET("/users", userHandler.GetUsers)
	api.POST("/users", userHandler.RegisterUser)
	router.Run("localhost:8080")
}
