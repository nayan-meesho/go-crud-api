package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"


	"github.com/gin-gonic/gin"
	"go-crud-api/models"
)

func CreateTask(c *gin.Context) {
	var t models.Task
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	created, err := models.AddTask(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	// Invalidate the cache
	models.RDB.Del(models.Ctx, "tasks")

	c.JSON(http.StatusOK, created)
}

func GetTasks(c *gin.Context) {

	// Try fetching from Redis
	cached, err := models.RDB.Get(models.Ctx, "tasks").Result()
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.String(http.StatusOK, cached)
		return
	}

	// Not in Redis — fetch from DB
	tasks, err := models.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	// Marshal and Cache the tasks in Redis
	jsonTasks, err := json.Marshal(tasks)
	if err == nil {
		// cache for 30 seconds
		models.RDB.Set(models.Ctx, "tasks", jsonTasks, 30*time.Second)
	}

	c.JSON(http.StatusOK, tasks)
}

func UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var t models.Task
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updated, ok, err := models.UpdateTask(id, t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Invalidate the cache
	models.RDB.Del(models.Ctx, "tasks")

	c.JSON(http.StatusOK, updated)
}

func DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	ok, err := models.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Invalidate the cache
	models.RDB.Del(models.Ctx, "tasks")

	c.Status(http.StatusNoContent)
}
