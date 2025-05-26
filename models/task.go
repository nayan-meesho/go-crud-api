package models

import (
    "log"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

// Task represents a to-do task with an ID, Title, and Done status.
type Task struct {
    ID      uint   `gorm:"primaryKey" json:"id"`
    Title   string `json:"title"`
    Done    bool   `json:"done"`
}

// InitDB initializes the DB connection
func InitDB(dsn string) {
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database:  %v", err)
    }

    if err = DB.AutoMigrate(&Task{}); err != nil {
        log.Fatalf("Failed to migrate database schema: %v", err)
    }

    log.Println("Connected to MySQL using GORM!")
}

// AddTask adds a new task to the in-memory list and assigns it an ID.
func AddTask(t Task) (Task, error) {
    if err := DB.Create(&t).Error; err != nil {
		return Task{}, err
	}
	return t, nil
}

func GetAllTasks() ([]Task, error) {
    var tasks []Task
	if err := DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func UpdateTask(id int, updated Task) (Task, bool, error) {
    var task Task
    if err := DB.First(&task, id).Error; err != nil {
		return Task{}, false, nil
	}

	task.Title = updated.Title
	task.Done = updated.Done

	if err := DB.Save(&task).Error; err != nil {
		return Task{}, false, err
	}

	return task, true, nil
}

func DeleteTask(id int) (bool, error) {
    res := DB.Delete(&Task{}, id)

    if res.Error != nil {
        return false, res.Error
    }

    return res.RowsAffected > 0, nil
}
