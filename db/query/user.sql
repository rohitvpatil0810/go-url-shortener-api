-- name: CreateUser :one
INSERT INTO "Users" (
  username, password_hash
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM "Users"
WHERE username = $1;

-- name: GetUserById :one
SELECT * FROM "Users"
WHERE id = $1;

-- name: GetUsers :many
SELECT * FROM "Users" ORDER BY created_at LIMIT $1 OFFSET $2;

-- name: DeleteUserById :exec 
DELETE FROM "Users" WHERE id = $1;

-- name: DeleteAllUsers :exec
DELETE FROM "Users";