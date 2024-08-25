package validations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUrlRequestValidations(t *testing.T) {
	createUrlRequestData := CreateUrlRequestParams{
		Url:    "https://www.google.com",
		UserID: "123e4567-e89b-12d3-a456-426614174000",
	}

	err := Validate.Struct(createUrlRequestData)
	assert.NoError(t, err, "Error should not be nil")

	badCreateUrlRequestData := CreateUrlRequestParams{
		Url:    "google.com",
		UserID: "123e4567-426614174000",
	}

	err = Validate.Struct(badCreateUrlRequestData)
	assert.Error(t, err, "Error should not be nil")
}

func TestDeleteUrlsByIdRequestValidations(t *testing.T) {
	deleteUrlByIdParams := DeleteUrlByIdRequestParams{
		ID: "123e4567-e89b-12d3-a456-426614174000",
	}

	err := Validate.Struct(deleteUrlByIdParams)
	assert.NoError(t, err, "Error should be nil")

	badDeleteUrlByIdParams := DeleteUrlByIdRequestParams{
		ID: "123456-323493i9",
	}
	err = Validate.Struct(badDeleteUrlByIdParams)
	assert.Error(t, err, "Error should not be nil")
}

func TestDeleteUrlsByUserIdValidations(t *testing.T) {
	deleteUrlsByUserIdParams := DeleteUrlsByUserIdRequestParams{
		UserID: "123e4567-e89b-12d3-a456-426614174000",
	}

	err := Validate.Struct(deleteUrlsByUserIdParams)
	assert.NoError(t, err, "Error should be nil")

	badDeleteUrlsByUserIdParams := DeleteUrlsByUserIdRequestParams{
		UserID: "123456-323493i9",
	}
	err = Validate.Struct(badDeleteUrlsByUserIdParams)
	assert.Error(t, err, "Error should not be nil")
}

func TestGetUrlByIdValidations(t *testing.T) {
	getUrlByIdParams := GetUrlByIdRequestParams{
		ID: "123e4567-e89b-12d3-a456-426614174000",
	}

	err := Validate.Struct(getUrlByIdParams)
	assert.NoError(t, err, "Error should be nil")

	badGetUrlByIdParams := GetUrlByIdRequestParams{
		ID: "123456-323493i9",
	}
	err = Validate.Struct(badGetUrlByIdParams)
	assert.Error(t, err, "Error should not be nil")
}

func TestGetUrlByShortenedUrlValidations(t *testing.T) {
	getUrlByShortenedUrlParams := GetUrlByShortenedUrlRequestParams{
		ShortenedUrl: "aseDfa",
	}

	err := Validate.Struct(getUrlByShortenedUrlParams)
	assert.NoError(t, err, "Error should be nil")

	badGetUrlByShortenedUrlParams := GetUrlByShortenedUrlRequestParams{
		ShortenedUrl: "",
	}
	err = Validate.Struct(badGetUrlByShortenedUrlParams)
	assert.Error(t, err, "Error should not be nil")
}
