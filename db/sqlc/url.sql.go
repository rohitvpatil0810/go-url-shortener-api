// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: url.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createUrl = `-- name: CreateUrl :one
INSERT INTO "Urls" (
  shortened_url, original_url, user_id
) VALUES (
  $1, $2, $3
) RETURNING id, user_id, shortened_url, original_url, click_count, created_at
`

type CreateUrlParams struct {
	ShortenedUrl string        `json:"shortened_url"`
	OriginalUrl  string        `json:"original_url"`
	UserID       uuid.NullUUID `json:"user_id"`
}

func (q *Queries) CreateUrl(ctx context.Context, arg CreateUrlParams) (Url, error) {
	row := q.db.QueryRowContext(ctx, createUrl, arg.ShortenedUrl, arg.OriginalUrl, arg.UserID)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ShortenedUrl,
		&i.OriginalUrl,
		&i.ClickCount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAllUrls = `-- name: DeleteAllUrls :exec
DELETE FROM "Urls"
`

func (q *Queries) DeleteAllUrls(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllUrls)
	return err
}

const deleteUrlById = `-- name: DeleteUrlById :exec
DELETE FROM "Urls"
WHERE id = $1
`

func (q *Queries) DeleteUrlById(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUrlById, id)
	return err
}

const deleteUrlsByUserId = `-- name: DeleteUrlsByUserId :exec
DELETE FROM "Urls"
WHERE user_id = $1
`

func (q *Queries) DeleteUrlsByUserId(ctx context.Context, userID uuid.NullUUID) error {
	_, err := q.db.ExecContext(ctx, deleteUrlsByUserId, userID)
	return err
}

const getUrlById = `-- name: GetUrlById :one
SELECT id, user_id, shortened_url, original_url, click_count, created_at FROM "Urls"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUrlById(ctx context.Context, id uuid.UUID) (Url, error) {
	row := q.db.QueryRowContext(ctx, getUrlById, id)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ShortenedUrl,
		&i.OriginalUrl,
		&i.ClickCount,
		&i.CreatedAt,
	)
	return i, err
}

const getUrlByShortenedUrl = `-- name: GetUrlByShortenedUrl :one
SELECT id, user_id, shortened_url, original_url, click_count, created_at FROM "Urls"
WHERE shortened_url = $1 LIMIT 1
`

func (q *Queries) GetUrlByShortenedUrl(ctx context.Context, shortenedUrl string) (Url, error) {
	row := q.db.QueryRowContext(ctx, getUrlByShortenedUrl, shortenedUrl)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ShortenedUrl,
		&i.OriginalUrl,
		&i.ClickCount,
		&i.CreatedAt,
	)
	return i, err
}

const getUrlsByUserId = `-- name: GetUrlsByUserId :many
SELECT id, user_id, shortened_url, original_url, click_count, created_at FROM "Urls"
WHERE user_id = $1 ORDER BY created_at LIMIT $2 OFFSET $3
`

type GetUrlsByUserIdParams struct {
	UserID uuid.NullUUID `json:"user_id"`
	Limit  int32         `json:"limit"`
	Offset int32         `json:"offset"`
}

func (q *Queries) GetUrlsByUserId(ctx context.Context, arg GetUrlsByUserIdParams) ([]Url, error) {
	rows, err := q.db.QueryContext(ctx, getUrlsByUserId, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Url
	for rows.Next() {
		var i Url
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ShortenedUrl,
			&i.OriginalUrl,
			&i.ClickCount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const incrementClickCount = `-- name: IncrementClickCount :exec
UPDATE "Urls"
SET click_count = click_count + 1
WHERE id = $1
`

func (q *Queries) IncrementClickCount(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, incrementClickCount, id)
	return err
}
