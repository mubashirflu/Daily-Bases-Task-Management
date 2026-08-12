package main

import (
	"task-management/controllers"
	"task-management/database"
	"task-management/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectDataBase()

	database.DB.AutoMigrate(
		&models.User{},
		&models.Task{},
	)

	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := router.Group("/api")

	// Auth
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	// Tasks
	api.POST("/tasks", controllers.CreateTask)
	api.GET("/tasks", controllers.GetDataTask)
	api.GET("/tasks/:id", controllers.GetTask)
	api.PUT("/tasks/:id", controllers.UpdateTask)
	api.DELETE("/tasks/:id", controllers.DeleteTask)

	router.Run(":8080")
}
