-- name: CreateUrl :one
INSERT INTO "Urls" (
  shortened_url, original_url, user_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUrls :many
SELECT * FROM "Urls" ORDER BY created_at LIMIT $1 OFFSET $2;

-- name: GetUrlByShortenedUrl :one
SELECT * FROM "Urls"
WHERE shortened_url = $1 LIMIT 1;

-- name: GetUrlById :one
SELECT * FROM "Urls"
WHERE id = $1 LIMIT 1;

-- name: GetUrlsByUserId :many
SELECT * FROM "Urls"
WHERE user_id = $1 ORDER BY created_at LIMIT $2 OFFSET $3;

-- name: IncrementClickCount :exec
UPDATE "Urls"
SET click_count = click_count + 1
WHERE id = $1;

-- name: DeleteUrlById :exec
DELETE FROM "Urls"
WHERE id = $1;

-- name: DeleteUrlsByUserId :exec
DELETE FROM "Urls"
WHERE user_id = $1;

-- name: DeleteAllUrls :exec
DELETE FROM "Urls";