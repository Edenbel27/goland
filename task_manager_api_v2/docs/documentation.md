
# 📝 Task Manager API v2 Documentation

This version of the Task Manager API introduces **MongoDB** integration for persistent task storage. Authentication is not included in this version.

---

## 🚀 Technologies Used
- Go (Golang)
- Gin Web Framework
- MongoDB

---

## ⚙️ Project Setup

1. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

2. **Ensure MongoDB is Running Locally** on port `27017`.

3. **Run the Application**:
   ```bash
   go run main.go
   ```

---

## ✅ API Endpoints

### GET `/tasks`
Fetch all tasks stored in the database.

**Response:**
```json
[
  {
    "id": "61234abc...",
    "title": "Sample Task",
    "description": "This is a task",
    "status": "pending"
  }
]
```

---

### GET `/tasks/:id`
Fetch a single task by its MongoDB ObjectID.

---

### POST `/tasks`
Create a new task.

**Request Body:**
```json
{
  "title": "New Task",
  "description": "Describe the task",
  "status": "pending"
}
```

---

### PUT `/tasks/:id`
Update an existing task.

**Request Body:**
```json
{
  "title": "Updated Task",
  "description": "Updated Description",
  "status": "completed"
}
```

---

### DELETE `/tasks/:id`
Delete a task using its ID.

---

## 🧩 MongoDB Task Model

```go
type Task struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    Title       string             `bson:"title"`
    Description string             `bson:"description"`
    Status      string             `bson:"status"`
}
```

---

## 🧪 Testing Tips

Use tools like [Postman](https://www.postman.com/) or `curl` to test your endpoints.

---