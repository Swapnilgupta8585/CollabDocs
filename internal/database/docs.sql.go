// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: docs.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createDoc = `-- name: CreateDoc :one
INSERT INTO docs(
    id,
    created_at,
    updated_at,
    user_id
)VALUES(
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1
)
RETURNING id, created_at, updated_at, user_id, content
`

func (q *Queries) CreateDoc(ctx context.Context, userID uuid.UUID) (Doc, error) {
	row := q.db.QueryRowContext(ctx, createDoc, userID)
	var i Doc
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.Content,
	)
	return i, err
}

const deleteDocByID = `-- name: DeleteDocByID :exec
DELETE FROM docs
WHERE id=$1
`

func (q *Queries) DeleteDocByID(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteDocByID, id)
	return err
}

const getDocByID = `-- name: GetDocByID :one
SELECT id, created_at, updated_at, user_id, content FROM docs
WHERE id=$1
`

func (q *Queries) GetDocByID(ctx context.Context, id uuid.UUID) (Doc, error) {
	row := q.db.QueryRowContext(ctx, getDocByID, id)
	var i Doc
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.Content,
	)
	return i, err
}

const updateContent = `-- name: UpdateContent :exec
UPDATE docs
SET content=$1, updated_at=NOW()
WHERE id=$2
`

type UpdateContentParams struct {
	Content string
	ID      uuid.UUID
}

func (q *Queries) UpdateContent(ctx context.Context, arg UpdateContentParams) error {
	_, err := q.db.ExecContext(ctx, updateContent, arg.Content, arg.ID)
	return err
}
