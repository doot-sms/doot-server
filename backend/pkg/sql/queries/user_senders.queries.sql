-- name: CreateUserSenderRelation :one
INSERT INTO
  user_senders (user_id, sender_id)
VALUES
  ($1, $2) RETURNING *;

-- name: DeleteUserSenderRelation :exec
DELETE FROM
  user_senders
WHERE
  id = $1;
