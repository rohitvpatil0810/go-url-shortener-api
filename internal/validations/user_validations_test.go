package validations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUserValidations(t *testing.T) {
	registerUserData := RegisterUserParams{
		Username: "us",
		Password: "pass",
	}

	err := Validate.Struct(registerUserData)
	assert.Error(t, err, "Error should not be nil")
}

func TestLoginUserValidations(t *testing.T) {
	loginUserData := LoginUserParams{
		Username: "us",
		Password: "pass",
	}

	err := Validate.Struct(loginUserData)
	assert.Error(t, err, "Error should not be nil")
}

func TestLogoutUserValidations(t *testing.T) {
	logoutUserData := LogoutUserParams{
		AccessToken: "",
	}

	err := Validate.Struct(logoutUserData)
	assert.Error(t, err, "Error should not be nil")
}
