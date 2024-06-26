// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: userquery.sql

package userstore

import (
	"context"

	"github.com/MWT-proger/go-sqlc-clean-architecture-example/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "auth"."user" (
          login, password, created_at
) VALUES (
  $1, $2, $3
)
RETURNING id, login, password, created_at
`

type CreateUserParams struct {
	Login     string
	Password  string
	CreatedAt pgtype.Timestamptz
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (models.AuthUser, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Login, arg.Password, arg.CreatedAt)
	var i models.AuthUser
	err := row.Scan(
		&i.ID,
		&i.Login,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "auth"."user"
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, login, password, created_at FROM "auth"."user"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (models.AuthUser, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i models.AuthUser
	err := row.Scan(
		&i.ID,
		&i.Login,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, login, password, created_at FROM "auth"."user"
ORDER BY login
`

func (q *Queries) ListUsers(ctx context.Context) ([]models.AuthUser, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.AuthUser
	for rows.Next() {
		var i models.AuthUser
		if err := rows.Scan(
			&i.ID,
			&i.Login,
			&i.Password,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
