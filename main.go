package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"

	"go-crud-api/kafka"
	"go-crud-api/middleware"
	"go-crud-api/models"
	"go-crud-api/routes"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Get env variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	redisAddr := os.Getenv("REDIS_ADDR")
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	port := os.Getenv("PORT")

	// Build DSN for MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	// Initialize DB and Redis
	models.InitDB(dsn)
	models.InitRedis(redisAddr)

	// Setup Gin router
	r := gin.Default()
	routes.SetupRoutes(r)

	// Add custom middleware
	r.Use(middleware.CustomLogger())

	// Init Kafka
	kafka.InitKafkaWriter(kafkaBroker, kafkaTopic)
	kafka.StartConsumer(kafkaBroker, kafkaTopic)

	// Start the server
	fmt.Printf("🚀 Server running at http://localhost:%s\n", port)
	log.Fatal(r.Run(":" + port))
}
