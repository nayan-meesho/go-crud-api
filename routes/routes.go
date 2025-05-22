package routes

import (
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "go-crud-api/handlers"
)

func SetupRoutes() *chi.Mux {
    r := chi.NewRouter()

    // Built-in logging middleware
    r.Use(middleware.Logger)
    r.Route("/tasks", func(r chi.Router) {
        r.Get("/", handlers.GetTasks)          // GET /tasks
        r.Post("/", handlers.CreateTask)       // POST /tasks
        r.Put("/{id}", handlers.UpdateTask)    // PUT /tasks/{id}
        r.Delete("/{id}", handlers.DeleteTask) // DELETE /tasks/{id}
    })

    return r
}
