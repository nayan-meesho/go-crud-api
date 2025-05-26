package routes

import (
    "go-crud-api/handlers"

    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {


    tasks := r.Group("/tasks") 
    {
        tasks.GET("/", handlers.GetTasks)
        tasks.POST("/", handlers.CreateTask)
        tasks.PUT("/:id", handlers.UpdateTask)
        tasks.DELETE("/:id", handlers.DeleteTask)
    }
    
}
