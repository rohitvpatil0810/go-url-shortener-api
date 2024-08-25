package validations

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func init() {
	Validate = validator.New()

	// user validations
	Validate.RegisterStructValidationMapRules(registerUserRules, RegisterUserParams{})
	Validate.RegisterStructValidationMapRules(loginUserRules, LoginUserParams{})
	Validate.RegisterStructValidationMapRules(logoutUserRules, LogoutUserParams{})

	// url validations
	Validate.RegisterStructValidationMapRules(createUrlRequestRules, CreateUrlRequestParams{})
	Validate.RegisterStructValidationMapRules(deleteUrlByIdRequestRules, DeleteUrlByIdRequestParams{})
	Validate.RegisterStructValidationMapRules(deleteUrlsByUserIdRequestRules, DeleteUrlsByUserIdRequestParams{})
	Validate.RegisterStructValidationMapRules(getUrlByIdRequestRules, GetUrlByIdRequestParams{})
	Validate.RegisterStructValidationMapRules(getUrlByShortenedUrlRequestRules, GetUrlByShortenedUrlRequestParams{})
}
