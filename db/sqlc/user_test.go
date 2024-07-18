package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
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

	// delete user
	t.Cleanup(func() {
		id := uuid.UUID(user.ID)
		_ = testQueries.DeleteUserById(context.Background(), id)
	})
}
