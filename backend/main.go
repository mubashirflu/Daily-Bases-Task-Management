// package main

// import (
// 	"task-management/controllers"
// 	"task-management/database"
// 	"task-management/middleware"
// 	"task-management/models"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// )

// func main() {

// 	database.ConnectDataBase()

// 	database.DB.AutoMigrate(
// 		&models.User{},
// 		&models.Task{},
// 	)

// 	router := gin.Default()

// 	// CORS
// 	router.Use(cors.New(cors.Config{
// 		AllowOrigins:     []string{"http://localhost:5173"},
// 		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
// 		AllowCredentials: true,
// 	}))

// 	api := router.Group("/api")

// 	// =========================
// 	// PUBLIC ROUTES
// 	// =========================

// 	api.POST("/register", controllers.Register)
// 	api.POST("/login", controllers.Login)

// 	// =========================
// 	// PROTECTED ROUTES
// 	// =========================

// 	protected := api.Group("/")
// 	protected.Use(middleware.AuthMiddleware())

// 	protected.POST("/tasks", controllers.CreateTask)
// 	protected.GET("/tasks", controllers.GetDataTask)
// 	api.GET("/tasks/scheduled", controllers.GetScheduledTasks)
// 	protected.GET("/tasks/:id", controllers.GetTask)
// 	protected.PUT("/tasks/:id", controllers.UpdateTask)
// 	protected.DELETE("/tasks/:id", controllers.DeleteTask)

// 	// Start server
// 	router.Run(":8080")
// }

package main

import (
	"fmt"
	"log"

	"task-management/controllers"
	"task-management/database"
	"task-management/middleware"
	"task-management/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// =========================
	// DATABASE
	// =========================

	database.ConnectDataBase()

	// Check database
	var dbName string
	if err := database.DB.Raw(
		"SELECT current_database()",
	).Scan(&dbName).Error; err != nil {
		log.Fatal("Could not get database name:", err)
	}

	fmt.Println("DATABASE:", dbName)

	// Check schema
	var schema string
	if err := database.DB.Raw(
		"SELECT current_schema()",
	).Scan(&schema).Error; err != nil {
		log.Fatal("Could not get schema:", err)
	}

	fmt.Println("SCHEMA:", schema)

	// Check tasks columns
	var columns []string

	if err := database.DB.Raw(`
		SELECT column_name
		FROM information_schema.columns
		WHERE table_schema = 'public'
		AND table_name = 'tasks'
		ORDER BY ordinal_position
	`).Scan(&columns).Error; err != nil {
		log.Fatal("Could not get task columns:", err)
	}

	fmt.Println("TASK COLUMNS:", columns)

	// =========================
	// MIGRATION
	// =========================

	if err := database.DB.AutoMigrate(
		&models.User{},
		&models.Task{},
	); err != nil {
		log.Fatal("Migration failed:", err)
	}

	fmt.Println("Migration successful")

	// =========================
	// GIN ROUTER
	// =========================

	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
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
	protected.GET("/tasks/scheduled", controllers.GetScheduledTasks)
	protected.GET("/tasks/:id", controllers.GetTask)
	protected.PUT("/tasks/:id", controllers.UpdateTask)
	protected.DELETE("/tasks/:id", controllers.DeleteTask)

	// =========================
	// SERVER
	// =========================

	fmt.Println("Server running on http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
