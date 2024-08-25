package validations

type CreateUrlRequestParams struct {
	Url    string `json:"url"`
	UserID string `json:"user_id"`
}

var createUrlRequestRules = map[string]string{
	"Url":    "required,url",
	"UserID": "required,uuid",
}

type DeleteUrlByIdRequestParams struct {
	ID string `json:"id"`
}

var deleteUrlByIdRequestRules = map[string]string{
	"ID": "required,uuid",
}

type DeleteUrlsByUserIdRequestParams struct {
	UserID string `json:"user_id"`
}

var deleteUrlsByUserIdRequestRules = map[string]string{
	"UserID": "required,uuid",
}

type GetUrlByIdRequestParams struct {
	ID string `json:"id"`
}

var getUrlByIdRequestRules = map[string]string{
	"ID": "required,uuid",
}

type GetUrlByShortenedUrlRequestParams struct {
	ShortenedUrl string `json:"shortened_url"`
}

var getUrlByShortenedUrlRequestRules = map[string]string{
	"ShortenedUrl": "required",
}
