package main

import (
	"kode/database"
	"kode/handlers"
	"kode/middleware"
	"kode/repositories"
	"kode/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	authHandler := handlers.NewAuthHandler(userService)
	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/users", userHandler.GetAllUsers)
	}

	r.Run(":8080")
}
