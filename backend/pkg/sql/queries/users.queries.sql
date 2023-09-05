-- name: CreateUser :exec
INSERT INTO users (email, password)
VALUES ($1, $2);

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: UpdatePassword :exec
UPDATE users SET password = $1 WHERE id = $2;
