package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/rohitvpatil0810/go-url-shortener-api/internal/utils"
	"github.com/stretchr/testify/assert"
)

func clearUrlAndUserTable() {
	testQueries.DeleteAllUrls(context.Background())
	testQueries.DeleteAllUsers(context.Background())
}

func createRandomUrl(t *testing.T, user User) Url {
	arg := CreateUrlParams{
		ShortenedUrl: utils.RandomString(10),
		OriginalUrl:  utils.RandomUrl(),
		UserID:       user.ID,
	}

	url, err := testQueries.CreateUrl(context.Background(), arg)
	assert.NoError(t, err, "Error should be nil")

	assert.NotNil(t, url, "Url should not be nil")
	assert.Equal(t, arg.ShortenedUrl, url.ShortenedUrl, "ShortenedUrl should be same")
	assert.Equal(t, arg.OriginalUrl, url.OriginalUrl, "OriginalUrl should be same")
	assert.Equal(t, arg.UserID, url.UserID, "UserID should be same")
	assert.NotZero(t, url.ID, "ID should not be zero")

	return url
}

func TestCreateUrl(t *testing.T) {
	user := createRandomUser(t)
	createRandomUrl(t, user)

	t.Cleanup(clearUrlAndUserTable)
}

func TestDeleteAllUrls(t *testing.T) {
	user := createRandomUser(t)
	for i := 0; i < 5; i++ {
		createRandomUrl(t, user)
	}

	err := testQueries.DeleteAllUrls(context.Background())
	assert.NoError(t, err, "Error should be nil")

	args := GetUrlsParams{
		Limit:  5,
		Offset: 0,
	}

	urls, err := testQueries.GetUrls(context.Background(), args)
	assert.NoError(t, err, "Error should be nil")
	assert.Empty(t, urls, "Urls should be empty")

	t.Cleanup(clearUrlAndUserTable)
}

func TestDeleteUrlById(t *testing.T) {
	user := createRandomUser(t)
	url := createRandomUrl(t, user)

	err := testQueries.DeleteUrlById(context.Background(), url.ID)
	assert.NoError(t, err, "Error should be nil")

	url2, err := testQueries.GetUrlById(context.Background(), url.ID)
	assert.Error(t, err, "Error should not be nil")
	assert.ErrorIs(t, err, sql.ErrNoRows, "Error should be user.ErrNoRows")
	assert.Empty(t, url2, "Url should be empty")

	t.Cleanup(clearUrlAndUserTable)
}

func TestDeleteUrlsByUserId(t *testing.T) {
	user := createRandomUser(t)
	for i := 0; i < 5; i++ {
		createRandomUrl(t, user)
	}

	err := testQueries.DeleteUrlsByUserId(context.Background(), user.ID)
	assert.NoError(t, err, "Error should be nil")

	args := GetUrlsByUserIdParams{
		UserID: user.ID,
		Limit:  5,
		Offset: 0,
	}

	urls, err := testQueries.GetUrlsByUserId(context.Background(), args)
	assert.NoError(t, err, "Error should be nil")
	assert.Empty(t, urls, "Urls should be empty")

	t.Cleanup(clearUrlAndUserTable)
}

func TestGetUrlById(t *testing.T) {
	user := createRandomUser(t)
	url := createRandomUrl(t, user)

	url2, err := testQueries.GetUrlById(context.Background(), url.ID)
	assert.NoError(t, err, "Error should be nil")

	assert.NotNil(t, url2, "Url should not be nil")
	assert.Equal(t, url.ID, url2.ID, "ID should be same")
	assert.Equal(t, url.UserID, url2.UserID, "UserID should be same")
	assert.Equal(t, url.ShortenedUrl, url2.ShortenedUrl, "ShortenedUrl should be same")
	assert.Equal(t, url.OriginalUrl, url2.OriginalUrl, "OriginalUrl should be same")

	t.Cleanup(clearUrlAndUserTable)
}

func TestGetUrlByShortenedUrl(t *testing.T) {
	user := createRandomUser(t)
	url := createRandomUrl(t, user)

	url2, err := testQueries.GetUrlByShortenedUrl(context.Background(), url.ShortenedUrl)
	assert.NoError(t, err, "Error should be nil")

	assert.NotNil(t, url2, "Url should not be nil")
	assert.Equal(t, url.ID, url2.ID, "ID should be same")
	assert.Equal(t, url.UserID, url2.UserID, "UserID should be same")
	assert.Equal(t, url.ShortenedUrl, url2.ShortenedUrl, "ShortenedUrl should be same")
	assert.Equal(t, url.OriginalUrl, url2.OriginalUrl, "OriginalUrl should be same")

	t.Cleanup(clearUrlAndUserTable)
}

func TestGetUrls(t *testing.T) {
	user := createRandomUser(t)
	for i := 0; i < 5; i++ {
		createRandomUrl(t, user)
	}

	args := GetUrlsParams{
		Limit:  5,
		Offset: 0,
	}

	urls, err := testQueries.GetUrls(context.Background(), args)
	assert.NoError(t, err, "Error should be nil")
	assert.Len(t, urls, 5, "Length should be 5")

	for _, url := range urls {
		assert.NotNil(t, url, "Url should not be nil")
	}

	t.Cleanup(clearUrlAndUserTable)
}

func TestGetUrlsByUserId(t *testing.T) {
	user := createRandomUser(t)
	for i := 0; i < 5; i++ {
		createRandomUrl(t, user)
	}

	args := GetUrlsByUserIdParams{
		UserID: user.ID,
		Limit:  5,
		Offset: 0,
	}

	urls, err := testQueries.GetUrlsByUserId(context.Background(), args)
	assert.NoError(t, err, "Error should be nil")
	assert.Len(t, urls, 5, "Length should be 5")

	for _, url := range urls {
		assert.NotNil(t, url, "Url should not be nil")
	}

	t.Cleanup(clearUrlAndUserTable)
}

func TestIncrementClickCount(t *testing.T) {
	user := createRandomUser(t)
	url := createRandomUrl(t, user)

	err := testQueries.IncrementClickCount(context.Background(), url.ID)
	assert.NoError(t, err, "Error should be nil")

	url2, err := testQueries.GetUrlById(context.Background(), url.ID)
	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, url.ClickCount+1, url2.ClickCount, "ClickCount should be incremented by 1")

	t.Cleanup(clearUrlAndUserTable)
}
