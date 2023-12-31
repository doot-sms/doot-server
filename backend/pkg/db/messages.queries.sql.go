// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: messages.queries.sql

package db

import (
	"context"
	"database/sql"
)

const createMessage = `-- name: CreateMessage :one
INSERT INTO "messages"
  ("to", content, user_id, batch_id)
  VALUES ( $1, $2, $3, $4 )
RETURNING id, "to", content, batch_id, sent_at, created_at, updated_at, user_id, sender_id
`

type CreateMessageParams struct {
	To      string
	Content string
	UserID  int32
	BatchID sql.NullInt32
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) (Message, error) {
	row := q.db.QueryRowContext(ctx, createMessage,
		arg.To,
		arg.Content,
		arg.UserID,
		arg.BatchID,
	)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.To,
		&i.Content,
		&i.BatchID,
		&i.SentAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.SenderID,
	)
	return i, err
}

const getSenderMessages = `-- name: GetSenderMessages :many
SELECT messages.id, messages."to", messages.content, messages.batch_id, messages.sent_at, messages.created_at, messages.updated_at, messages.user_id, messages.sender_id
  FROM user_senders
    LEFT JOIN messages
    ON user_senders.user_id = messages.user_id
  WHERE messages.sent_at is null AND user_senders.sender_id = $1
`

type GetSenderMessagesRow struct {
	ID        sql.NullInt32
	To        sql.NullString
	Content   sql.NullString
	BatchID   sql.NullInt32
	SentAt    sql.NullTime
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	UserID    sql.NullInt32
	SenderID  sql.NullInt32
}

func (q *Queries) GetSenderMessages(ctx context.Context, senderID int32) ([]GetSenderMessagesRow, error) {
	rows, err := q.db.QueryContext(ctx, getSenderMessages, senderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetSenderMessagesRow
	for rows.Next() {
		var i GetSenderMessagesRow
		if err := rows.Scan(
			&i.ID,
			&i.To,
			&i.Content,
			&i.BatchID,
			&i.SentAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.SenderID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserMessages = `-- name: GetUserMessages :many
SELECT id, "to", content, batch_id, sent_at, created_at, updated_at, user_id, sender_id FROM messages WHERE user_id = $1
`

func (q *Queries) GetUserMessages(ctx context.Context, userID int32) ([]Message, error) {
	rows, err := q.db.QueryContext(ctx, getUserMessages, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.To,
			&i.Content,
			&i.BatchID,
			&i.SentAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.SenderID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const markAsSent = `-- name: MarkAsSent :one
UPDATE "messages"
  SET sent_at = now(), sender_id = $2
  WHERE id = $1
RETURNING id, "to", content, batch_id, sent_at, created_at, updated_at, user_id, sender_id
`

type MarkAsSentParams struct {
	ID       int32
	SenderID int32
}

func (q *Queries) MarkAsSent(ctx context.Context, arg MarkAsSentParams) (Message, error) {
	row := q.db.QueryRowContext(ctx, markAsSent, arg.ID, arg.SenderID)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.To,
		&i.Content,
		&i.BatchID,
		&i.SentAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.SenderID,
	)
	return i, err
}
