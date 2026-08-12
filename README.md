# Task Management System

A full-stack Task Management System built with **Vue.js** and **Go (Golang)**.
The application allows users to create, update, delete, and manage their tasks through a simple and responsive interface.

## 🚀 Live Demo


## 📌 Features

* Create new tasks
* View all tasks
* Update existing tasks
* Delete tasks
* Mark tasks as completed
* REST API integration
* Responsive user interface
* Frontend and backend separated architecture
* Git/GitHub based development workflow

## 🛠️ Technologies Used

### Frontend

* Vue.js
* JavaScript
* HTML5
* CSS3
* Vite

### Backend

* Go (Golang)
* REST API
* GORM

### Database

* PostgreSQL / MySQL

### Tools

* Git
* GitHub
* Vercel

## 📂 Project Structure

```text
task-management/
│
├── frontend/
│   ├── src/
│   ├── public/
│   ├── package.json
│   └── vite.config.js
│
├── backend/
│   ├── controllers/
│   ├── models/
│   ├── routes/
│   ├── main.go
│   └── go.mod
│
└── README.md
```

## ⚙️ Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/YOUR-USERNAME/YOUR-REPOSITORY.git
cd YOUR-REPOSITORY
```

## 💻 Frontend Setup

Go to the frontend directory:

```bash
cd frontend
```

Install dependencies:

```bash
npm install
```

Start the development server:

```bash
npm run dev
```

The frontend will normally run on:

```text
http://localhost:5173
```

## 🔧 Backend Setup

Go to the backend directory:

```bash
cd backend
```

Install Go dependencies:

```bash
go mod tidy
```

Run the backend:

```bash
go run main.go
```

The backend API will normally run on:

```text
http://localhost:8080
```

## 🔗 API Endpoints

| Method | Endpoint     | Description         |
| ------ | ------------ | ------------------- |
| GET    | `/tasks`     | Get all tasks       |
| GET    | `/tasks/:id` | Get a specific task |
| POST   | `/tasks`     | Create a new task   |
| PUT    | `/tasks/:id` | Update a task       |
| DELETE | `/tasks/:id` | Delete a task       |

## 🗄️ Database

The backend uses **GORM** for database interaction.

Example task structure:

```text
Task
├── ID
├── Title
├── Description
├── Completed
├── CreatedAt
└── UpdatedAt
```

## 🌐 Deployment

The frontend can be deployed using **Vercel**.

The backend can be deployed separately using a Go-compatible hosting platform.

After deployment, update the frontend API URL to point to the deployed backend.

Example:

```javascript
const API_URL = "https://your-backend-url.com";
```

## 📸 Screenshots

Add screenshots of your application here:

```text
screenshots/
├── dashboard.png
├── create-task.png
└── edit-task.png
```

## 🎯 Learning Goals

This project was created to practice:

* Vue.js frontend development
* Go backend development
* REST API development
* GORM
* Database integration
* Frontend-backend communication
* CRUD operations
* Git and GitHub
* Deployment

## 👨‍💻 Author

**Muhammad Mubashir**

Computer Science Student
Interested in Full-Stack Web Development

## 📄 License

This project is created for learning and portfolio purposes.
