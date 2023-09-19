-- name: CreateApiKey :one
INSERT INTO "user_api_keys"
(user_id, api_secret, expires_after)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteApiKey :exec
DELETE FROM "user_api_keys"
  WHERE api_key = $1;
