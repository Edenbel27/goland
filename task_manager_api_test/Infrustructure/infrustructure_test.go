package Infrustructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordService(t *testing.T) {
	svc := NewPasswordService()
	hashed, err := svc.HashPassword("password")
	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)
	match := svc.CompareHashedPassword(hashed, "password")
	assert.True(t, match)
}

func TestJWTService(t *testing.T) {
	svc := NewJWTService()
	token, err := svc.GenerateToken("user@example.com", "user")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

}
