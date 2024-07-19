package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func clearUserTable() {
	testQueries.DeleteAllUsers(context.Background())
}

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:     "user1",
		PasswordHash: "password",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	assert.NoError(t, err, "Error should be nil")

	assert.NotNil(t, user, "User should not be nil")
	assert.Equal(t, arg.Username, user.Username, "Username should be same")
	assert.Equal(t, arg.PasswordHash, user.PasswordHash, "PasswordHash should be same")
	assert.NotZero(t, user.ID, "ID should not be zero")

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)

	// clear user table
	t.Cleanup(clearUserTable)
}
