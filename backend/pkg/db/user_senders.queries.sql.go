// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: user_senders.queries.sql

package db

import (
	"context"
)

const createUserSenderRelation = `-- name: CreateUserSenderRelation :one
INSERT INTO user_senders (
  user_id, sender_id
) VALUES ( $1, $2 )
RETURNING id, user_id, sender_id, created_at, updated_at
`

type CreateUserSenderRelationParams struct {
	UserID   int32
	SenderID int32
}

func (q *Queries) CreateUserSenderRelation(ctx context.Context, arg CreateUserSenderRelationParams) (UserSender, error) {
	row := q.db.QueryRowContext(ctx, createUserSenderRelation, arg.UserID, arg.SenderID)
	var i UserSender
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.SenderID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUserSenderRelation = `-- name: DeleteUserSenderRelation :exec
DELETE FROM user_senders
WHERE id = $1
`

func (q *Queries) DeleteUserSenderRelation(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUserSenderRelation, id)
	return err
}