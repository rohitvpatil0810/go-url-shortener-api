// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"

	"github.com/google/uuid"
)

type Url struct {
	ID           uuid.UUID     `json:"id"`
	UserID       uuid.NullUUID `json:"user_id"`
	ShortenedUrl string        `json:"shortened_url"`
	OriginalUrl  string        `json:"original_url"`
	ClickCount   int64         `json:"click_count"`
	CreatedAt    sql.NullTime  `json:"created_at"`
}

type User struct {
	ID           uuid.UUID    `json:"id"`
	Username     string       `json:"username"`
	PasswordHash string       `json:"password_hash"`
	CreatedAt    sql.NullTime `json:"created_at"`
}
