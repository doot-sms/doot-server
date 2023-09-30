-- name: CreateRefreshToken :one
INSERT INTO
    refresh_tokens (user_id, user_agent)
VALUES
    ($ 1, $ 2) RETURNING *;