package validations

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func init() {
	Validate = validator.New()

	Validate.RegisterStructValidationMapRules(registerUserRules, RegisterUserParams{})
	Validate.RegisterStructValidationMapRules(loginUserRules, LoginUserParams{})
	Validate.RegisterStructValidationMapRules(logoutUserRules, LogoutUserParams{})
}
