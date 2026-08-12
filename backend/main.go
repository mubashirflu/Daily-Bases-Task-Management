package main

import (
	"task-management/controllers"
	"task-management/database"
	"task-management/middleware"
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

	// =========================
	// PUBLIC ROUTES
	// =========================

	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	// =========================
	// PROTECTED ROUTES
	// =========================

	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.POST("/tasks", controllers.CreateTask)
	protected.GET("/tasks", controllers.GetDataTask)
	protected.GET("/tasks/:id", controllers.GetTask)
	protected.PUT("/tasks/:id", controllers.UpdateTask)
	protected.DELETE("/tasks/:id", controllers.DeleteTask)

	// Start server
	router.Run(":8080")
}
