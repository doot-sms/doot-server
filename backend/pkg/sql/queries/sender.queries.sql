-- name: CreateSender :one
INSERT INTO
    senders (user_id, device_id)
VALUES
    ($ 1, $ 2) RETURNING *;

-- name: GetSenderByID :one
SELECT
    *
FROM
    senders
WHERE
    id = $ 1;

-- name: UpdateSender :one
UPDATE
    senders
SET
    user_id = $ 1,
    device_id = $ 2
WHERE
    id = $ 3 RETURNING *;