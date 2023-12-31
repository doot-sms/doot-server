// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: sender.queries.sql

package db

import (
	"context"
)

const createSender = `-- name: CreateSender :one
INSERT INTO senders (user_id, device_id)
VALUES ($1, $2)
RETURNING id, user_id, device_id, created_at, updated_at
`

type CreateSenderParams struct {
	UserID   int32
	DeviceID string
}

func (q *Queries) CreateSender(ctx context.Context, arg CreateSenderParams) (Sender, error) {
	row := q.db.QueryRowContext(ctx, createSender, arg.UserID, arg.DeviceID)
	var i Sender
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DeviceID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSenderByID = `-- name: GetSenderByID :one
SELECT id, user_id, device_id, created_at, updated_at FROM senders WHERE id = $1
`

func (q *Queries) GetSenderByID(ctx context.Context, id int32) (Sender, error) {
	row := q.db.QueryRowContext(ctx, getSenderByID, id)
	var i Sender
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DeviceID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateSender = `-- name: UpdateSender :one
UPDATE senders SET user_id = $1, device_id = $2 WHERE id = $3
RETURNING id, user_id, device_id, created_at, updated_at
`

type UpdateSenderParams struct {
	UserID   int32
	DeviceID string
	ID       int32
}

func (q *Queries) UpdateSender(ctx context.Context, arg UpdateSenderParams) (Sender, error) {
	row := q.db.QueryRowContext(ctx, updateSender, arg.UserID, arg.DeviceID, arg.ID)
	var i Sender
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DeviceID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
