
# 📝 Task Manager API v3 Documentation

This is the third version of the Task Manager API, enhanced to include **User Management**and **JWT Authentication**.

---

## 🚀 Technologies Used
- Go (Golang)
- Gin Web Framework
- MongoDB (via mgo or mongo-driver)
- JWT (JSON Web Tokens)

---

## ⚙️ Project Setup

1. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

2. **Set Environment Variables** in `.env`:
   ```env
   MONGO_URI=mongodb://localhost:27017
   JWT_SECRET=MySecret!
   ```

3. **Run the App**:
   ```bash
   go run main.go
   ```

---

## 🔐 Authentication

### POST `/register`
Registers a new user.

**Body:**
```json
{
  "username": "john_doe",
  "password": "secure123",
  "role": "admin"
}
```

---

### POST `/login`
Logs in the user and returns a JWT token.

**Body:**
```json
{
  "username": "john_doe",
  "password": "secure123"
}
```

**Response:**
```json
{
  "token": "jwt_token_here"
}
```

---

## 🔒 Protected Routes (JWT required)

### GET `/api/protected`
Returns a protected message if token is valid.

**Headers:**
```
Authorization: Bearer <token>
```

---

## 🛡️ Admin Routes

### GET `/api/admin/dashboard`
Accessible only to users with `admin` role.

---

## ✅ Task Management Endpoints

### GET `/tasks`
Fetch all tasks.

---

### GET `/tasks/:id`
Fetch task by ID.

---

### POST `/tasks`
Create a new task.

**Body:**
```json
{
  "title": "New Task",
  "description": "Task details",
  "status": "pending"
}
```

---

### PUT `/tasks/:id`
Update an existing task.

**Body:**
```json
{
  "title": "Updated Task",
  "description": "Updated details",
  "status": "completed"
}
```

---

### DELETE `/tasks/:id`
Delete a task by ID.

---

## 🧩 MongoDB Models

### User
```go
type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty"`
    Username string             `bson:"username"`
    Password string             `bson:"password"`
    Role     string             `bson:"role"`
}
```

### Task
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

Use [Postman](https://www.postman.com/) or `curl` to test the endpoints. Remember to include the JWT token in the `Authorization` header for protected routes.

---
