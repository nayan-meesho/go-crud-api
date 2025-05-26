package main

import (
	"fmt"
	"log"

	"go-crud-api/models"
	"go-crud-api/routes"
	"go-crud-api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the DB connection
	dsn := "root:Nayan123!@#@tcp(127.0.0.1:3306)/tasks_db"
	models.InitDB(dsn)


	// Setup routes
	r := gin.Default()
	routes.SetupRoutes(r)

	// Use the custom logger middleware
	r.Use(middleware.CustomLogger())

	// Start the server
	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(r.Run(":8080"))
}
