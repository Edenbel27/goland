package Domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserFields(t *testing.T) {
	user := User{Email: "test@example.com", Password: "pass", Role: "admin"}
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, "pass", user.Password)
	assert.Equal(t, "admin", user.Role)
}

func TestTaskFields(t *testing.T) {
	task := Task{Title: "title", Description: "desc", Status: "pending"}
	assert.Equal(t, "title", task.Title)
	assert.Equal(t, "desc", task.Description)
	assert.Equal(t, "pending", task.Status)
}
