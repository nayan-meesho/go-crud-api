package main

import (
    "fmt"
    "log"
    "net/http"

    "go-crud-api/routes"
    "go-crud-api/models"
)

func main() {
    // Initialize the DB connection
    dsn := "root:Nayan123!@#@tcp(127.0.0.1:3306)/tasks_db"

    // Initialize the DB connection
    models.InitDB(dsn)

    // Setup routes
    r := routes.SetupRoutes()

    // Start the server
    fmt.Println("🚀 Server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
