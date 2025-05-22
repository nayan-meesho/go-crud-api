# Go CRUD API 🧩

A simple CRUD (Create, Read, Update, Delete) API built using **Go** and the **chi router**. This project manages tasks in memory and demonstrates modular Go web development.

---


## ✨ Features

- ✅ Add a new task
- 📋 Get all tasks
- 📝 Update a task by ID
- ❌ Delete a task by ID
- 🧱 Modular folder structure
- 🛡️ Middleware for logging

---


## 📁 Project Structure

go-crud-api/
├── handlers/ # HTTP handler functions (Create, Read, Update, Delete)
│ └── task.go
├── middleware/ # Custom middlewares (e.g., logging)
│ └── logger.go
├── models/ # Data models and logic
│ └── task.go
├── routes/ # Route definitions
│ └── routes.go
├── main.go # Entry point
└── go.mod # Module definition

---


## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/your-username/go-crud-api.git
cd go-crud-api
```


### 2. Run the app

```bash
go run main.go
```

#### Server starts on:

```bash
http://localhost:8080
```


## 🧪 API Endpoints

| Method | Endpoint      | Description         |
| ------ | ------------- | ------------------- |
| POST   | `/tasks`      | Create a new task   |
| GET    | `/tasks`      | Get all tasks       |
| PUT    | `/tasks/{id}` | Update a task by ID |
| DELETE | `/tasks/{id}` | Delete a task by ID |


## 🧱 Tech Stack

Go (Golang)

chi router (github.com/go-chi/chi)

net/http standard library
