-- name: GetUSSRRequestById :one
SELECT * FROM user_sender_reqs where id = $1;

-- name: GetUSSRRequestsByUserIdSenderId :one
SELECT
  *
FROM
  user_sender_reqs
WHERE
  user_id = $1
  and sender_id = $2
  and "status" = $3
LIMIT 1;

-- name: GetPendingRequestsByUserId :many
SELECT * FROM user_sender_reqs WHERE user_id = $1 and "status" = $2;

-- name: GetPendingRequestsBySenderId :many
SELECT * FROM user_sender_reqs WHERE sender_id = $1 and "status" = $2;

-- name: CreateUSSRRequest :one
INSERT INTO
  user_sender_reqs (user_id, sender_id, requestor)
VALUES
  ($1, $2, $3) RETURNING *;

-- name: RejectRequest :one
UPDATE
  user_sender_reqs
SET
  "status" = 'rejected'
WHERE
  id = $1 RETURNING *;

-- name: AcceptRequest :one
UPDATE
  user_sender_reqs
SET
  "status" = 'accepted'
WHERE
  id = $1 RETURNING *;
