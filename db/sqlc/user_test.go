package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/rohitvpatil0810/go-url-shortener-api/internal/utils"
	"github.com/stretchr/testify/assert"
)

func clearUserTable() {
	testQueries.DeleteAllUsers(context.Background())
}

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:     utils.RandomUsername(),
		PasswordHash: utils.RandomPassword(),
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

func TestGetUserById(t *testing.T) {
	user := createRandomUser(t)

	user2, err := testQueries.GetUserById(context.Background(), user.ID)
	assert.NoError(t, err, "Error should be nil")

	assert.NotNil(t, user2, "User should not be nil")
	assert.Equal(t, user.ID, user2.ID, "ID should be same")
	assert.Equal(t, user.Username, user2.Username, "Username should be same")
	assert.Equal(t, user.PasswordHash, user2.PasswordHash, "PasswordHash should be same")

	// clear user table
	t.Cleanup(clearUserTable)
}

func TestGetUserByUsername(t *testing.T) {
	user := createRandomUser(t)

	user2, err := testQueries.GetUserByUsername(context.Background(), user.Username)
	assert.NoError(t, err, "Error should be nil")

	assert.NotNil(t, user2, "User should not be nil")
	assert.Equal(t, user.ID, user2.ID, "ID should be same")
	assert.Equal(t, user.Username, user2.Username, "Username should be same")
	assert.Equal(t, user.PasswordHash, user2.PasswordHash, "PasswordHash should be same")

	// clear user table
	t.Cleanup(clearUserTable)
}

func TestDeleteUserById(t *testing.T) {
	user := createRandomUser(t)
	fmt.Println(user)
	err := testQueries.DeleteUserById(context.Background(), user.ID)
	assert.NoError(t, err, "Error should be nil")

	user2, err := testQueries.GetUserById(context.Background(), user.ID)
	fmt.Print(user2)
	assert.Error(t, err, "Error should not be nil")
	assert.ErrorIs(t, err, sql.ErrNoRows, "Error should be sql.ErrNoRows")
	assert.Empty(t, user2, "User should be empty")

	// clear user table
	t.Cleanup(clearUserTable)
}

func TestGetUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	args := GetUsersParams{
		Offset: 0,
		Limit:  5,
	}

	users, err := testQueries.GetUsers(context.Background(), args)
	assert.NoError(t, err, "Error should be nil")
	assert.Len(t, users, 5, "Length should be 5")

	for _, user := range users {
		assert.NotNil(t, user, "User should not be nil")
	}

	// clear user table
	t.Cleanup(clearUserTable)
}

func TestDeleteAllUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	err := testQueries.DeleteAllUsers(context.Background())
	assert.NoError(t, err, "Error should be nil")

	args := GetUsersParams{
		Offset: 0,
		Limit:  5,
	}

	users, err := testQueries.GetUsers(context.Background(), args)
	assert.NoError(t, err, "Error should be nil")
	assert.Len(t, users, 0, "Length should be 0")
}
