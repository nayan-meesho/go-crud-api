package main

import (
	"fmt"
	"log"

	"go-crud-api/middleware"
	"go-crud-api/models"
	"go-crud-api/routes"

	"go-crud-api/kafka"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the DB connection
	dsn := "root:Nayan123!@#@tcp(127.0.0.1:3306)/tasks_db"
	models.InitDB(dsn)
	models.InitRedis("localhost:6379")

	// Setup routes
	r := gin.Default()
	routes.SetupRoutes(r)

	// Use the custom logger middleware
	r.Use(middleware.CustomLogger())

	// Initialize Kafka producer and consumer
	kafka.InitKafkaWriter("localhost:9092", "task-notifications")
	kafka.StartConsumer("localhost:9092", "task-notifications")

	// Start the server
	fmt.Println("🚀 Server running at http://localhost:8081")
	log.Fatal(r.Run(":8081"))
}
