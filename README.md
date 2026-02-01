# Golang Notes Management API

A simple **Notes Management REST API** built with **Golang**, using a **structured directory** (handlers, services, repositories, models, database) to demonstrate clean architecture.  
This project is perfect as a **portfolio project** for backend development.

---

## 🔹 Features

- Create, Read, Update, Delete (CRUD) notes
- Structured code: `handlers`, `services`, `repositories`, `models`, `database`
- Configurable via `.env`
- Ready for deployment (Railway, Heroku, etc.)

---

## 🚀 Installation & Run Locally

1. Clone the repository:

```bash
git clone https://github.com/L4st4ken/Golang-Notes-Management-with-Structured-Directory.git
cd Golang-Notes-Management-with-Structured-Directory
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the API::
```bash
go run ./cmd/api
```

## 📝 API Endpoints

| Method | Endpoint       | Description           |
|--------|----------------|---------------------|
| GET    | /test          | Test server status  |
| GET    | /api/note      | Get all notes       |
| GET    | /api/note/:id  | Get a note by ID    |
| POST   | /api/note      | Create a new note   |
| PUT    | /api/note/:id  | Update a note by ID |
| DELETE | /api/note/:id  | Delete a note by ID |
