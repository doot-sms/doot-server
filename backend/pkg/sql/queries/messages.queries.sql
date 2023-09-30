-- name: CreateMessage :one
INSERT INTO
  "messages" ("to", content, user_id, batch_id)
VALUES
  ($ 1, $ 2, $ 3, $ 4) RETURNING *;

-- name: MarkAsSent :one
UPDATE
  "messages"
SET
  sent_at = now(),
  sender_id = $ 2
WHERE
  id = $ 1 RETURNING *;

-- name: GetUserMessages :many
SELECT
  *
FROM
  messages
WHERE
  user_id = $ 1;

-- name: GetSenderMessages :many
SELECT
  messages.*
FROM
  user_senders
  LEFT JOIN messages ON user_senders.user_id = messages.user_id
WHERE
  messages.sent_at is null
  AND user_senders.sender_id = $ 1;