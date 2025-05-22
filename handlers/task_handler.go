package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    "go-crud-api/models"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var t models.Task
	err := json.NewDecoder(r.Body).Decode(&t);
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    created := models.AddTask(t)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(created)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(models.GetAllTasks())
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    var t models.Task
    if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    updated, ok := models.UpdateTask(id, t)
    if !ok {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updated)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    ok := models.DeleteTask(id)
    if !ok {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
