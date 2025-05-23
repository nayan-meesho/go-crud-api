package models

import (
    "database/sql"
    // "errors"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Task represents a to-do task with an ID, Title, and Done status.
type Task struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Done    bool   `json:"done"`
}

// InitDB initializes the DB connection
func InitDB(dataSourceName string) {
    var err error
    db, err = sql.Open("mysql", dataSourceName)
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }

    if err = db.Ping(); err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

    log.Println("Connected to MySQL database!")
}

// AddTask adds a new task to the in-memory list and assigns it an ID.
func AddTask(t Task) (Task, error) {
    result, err := db.Exec("INSERT INTO tasks (title, done) VALUES (?, ?)", t.Title, t.Done)
    if err != nil {
        return Task{}, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return Task{}, err
    }

    t.ID = int(id)
    return t, nil
}

func GetAllTasks() ([]Task, error) {
    rows, err := db.Query("SELECT id, title, done FROM tasks")

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    var tasks []Task

    for rows.Next() {
        var t Task
        var doneInt int
        if err := rows.Scan(&t.ID, &t.Title, &doneInt); err != nil {
            return nil, err
        }
        t.Done = doneInt == 1
        tasks = append(tasks, t)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return tasks, nil
}

func UpdateTask(id int, updated Task) (Task, bool, error) {
    res, err := db.Exec("UPDATE tasks SET title=?, done=? WHERE id=?", updated.Title, updated.Done, id)
    if err != nil {
        return Task{}, false, err
    }

    rowsAffected, err := res.RowsAffected()
    if err != nil {
        return Task{}, false, err
    }

    if rowsAffected == 0 {
        return Task{}, false, nil
    }
    updated.ID = id
    return updated, true, nil
}

func DeleteTask(id int) (bool, error) {
    res, err := db.Exec("DELETE FROM tasks WHERE id=?", id)
    if err != nil {
        return false, err
    }

    rowsAffected, err := res.RowsAffected()
    if err != nil {
        return false, err
    }

    return rowsAffected > 0, nil
}
