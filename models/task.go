package models

// Task represents a to-do task with an ID, Title, and Done status.
type Task struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Done    bool   `json:"done"`
}

var Tasks []Task
var nextID = 1

// AddTask adds a new task to the in-memory list and assigns it an ID.
func AddTask(t Task) Task {
    t.ID = nextID
    nextID++
    Tasks = append(Tasks, t)
    return t
}

func GetAllTasks() []Task {
    return Tasks
}

func UpdateTask(id int, updated Task) (Task, bool) {
    for i, t := range Tasks {
        if t.ID == id {
            Tasks[i].Title = updated.Title
            Tasks[i].Done = updated.Done
            return Tasks[i], true
        }
    }
    return Task{}, false
}

func DeleteTask(id int) bool {
    for i, t := range Tasks {
        if t.ID == id {
            Tasks = append(Tasks[:i], Tasks[i+1:]...)
            return true
        }
    }
    return false
}
