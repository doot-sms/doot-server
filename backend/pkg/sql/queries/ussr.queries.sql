-- name: GetUSSRRequestsByUserIdSenderId :many
SELECT
  user_id,
  sender_id,
  requestor,
  "status"
FROM
  user_sender_reqs
where
  user_id = $ 1
  and sender_id = $ 2
  and "status" = 3;

-- name: CreateUSSRRequest :one
INSERT INTO
  user_sender_reqs (user_id, sender_id, requestor)
VALUES
  ($ 1, $ 2, $ 3) RETURNING *;

-- name: RejectRequest :one
UPDATE
  user_sender_reqs
SET
  "status" = 'rejected'
WHERE
  id = $ 1 RETURNING *;

-- name: AcceptRequest :one
UPDATE
  user_sender_reqs
SET
  status = 'accepted'
WHERE
  id = $ 1 RETURNING *;