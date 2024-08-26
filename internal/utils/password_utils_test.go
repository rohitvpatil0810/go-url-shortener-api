package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err, "error should be nil")
	assert.NotEmpty(t, hashedPassword, "hashedPassword should not be empty")
}

func TestComparePassword(t *testing.T) {
	password := "password"
	hashedPassword, _ := HashPassword(password)
	result := ComparePassword(password, hashedPassword)
	assert.True(t, result, "passwords should match")
}
