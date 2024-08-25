package validations

type RegisterUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var registerUserRules = map[string]string{
	"Username": "required,alphanum,min=3,max=40",
	"Password": "required,min=8,max=40",
}

type LoginUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var loginUserRules = map[string]string{
	"Username": "required,alphanum,min=3,max=40",
	"Password": "required,min=8,max=40",
}

type LogoutUserParams struct {
	AccessToken string `json:"access_token"`
}

var logoutUserRules = map[string]string{
	"AccessToken": "required",
}
