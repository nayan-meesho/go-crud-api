package main

import (
    "fmt"
    "net/http"
    "go-crud-api/routes"
)

func main() {
    r := routes.SetupRoutes()
    fmt.Println("🚀 Server running at http://localhost:8080")
    http.ListenAndServe(":8080", r)
}
