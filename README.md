# Tally - Task Management App

A full-stack Task Management application built with **Vue.js** and **Golang**.

Users can register, login securely, receive a JWT token, and manage their tasks.

---

## 🚀 Features

### Authentication

- User registration
- Password hashing using bcrypt
- Secure login
- JWT authentication
- Protected API routes
- Token-based authorization

### Task Management

- Create tasks
- View tasks
- Update tasks
- Mark tasks as completed/pending
- Delete tasks
- Task progress percentage
- Completed task counter

### Security

- Passwords are never stored as plain text
- Passwords are hashed using bcrypt
- JWT tokens are used for authentication
- Protected task APIs require an Authorization token
- CORS configured for frontend

---

## 🛠️ Tech Stack

### Frontend

- Vue.js
- JavaScript
- HTML
- CSS
- Fetch API

### Backend

- Golang
- Gin
- GORM
- JWT
- bcrypt

### Database

- PostgreSQL

---

## 📁 Project Structure

```text
task-management/
│
├── backend/
│   ├── controllers/
│   │   ├── auth.go
│   │   └── task.go
│   │
│   ├── database/
│   │   └── database.go
│   │
│   ├── middleware/
│   │   └── auth.go
│   │
│   ├── models/
│   │   ├── user.go
│   │   └── task.go
│   │
│   ├── utils/
│   │   └── jwt.go
│   │
│   ├── main.go
│   ├── go.mod
│   └── go.sum
│
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   ├── views/
│   │   └── ...
│   │
│   ├── package.json
│   └── ...
│
└── README.md
