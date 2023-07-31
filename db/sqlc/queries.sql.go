// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: queries.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createNote = `-- name: CreateNote :one
INSERT INTO notes (id, title, username, text, created_at, updated_at)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING id
`

type CreateNoteParams struct {
	ID        uuid.UUID
	Title     string
	Username  string
	Text      sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateNote(ctx context.Context, arg *CreateNoteParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createNote,
		arg.ID,
		arg.Title,
		arg.Username,
		arg.Text,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const deleteNote = `-- name: DeleteNote :one
DELETE 
FROM notes
WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteNote(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, deleteNote, id)
	err := row.Scan(&id)
	return id, err
}

const getAllNotesFromUser = `-- name: GetAllNotesFromUser :many
SELECT id, title, username, text, created_at, updated_at
FROM notes
WHERE username = $1
`

func (q *Queries) GetAllNotesFromUser(ctx context.Context, username string) ([]Note, error) {
	rows, err := q.db.QueryContext(ctx, getAllNotesFromUser, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Note{}
	for rows.Next() {
		var i Note
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Username,
			&i.Text,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getUser = `-- name: GetUser :one
SELECT username, password, email FROM users
WHERE username = $1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(&i.Username, &i.Password, &i.Email)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT username, password, email
FROM users
ORDER BY username
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.Username, &i.Password, &i.Email); err != nil {
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

const registerUser = `-- name: RegisterUser :one
INSERT INTO users (username, password, email)
VALUES ($1,$2,$3)
RETURNING username
`

type RegisterUserParams struct {
	Username string
	Password string
	Email    string
}

func (q *Queries) RegisterUser(ctx context.Context, arg *RegisterUserParams) (string, error) {
	row := q.db.QueryRowContext(ctx, registerUser, arg.Username, arg.Password, arg.Email)
	var username string
	err := row.Scan(&username)
	return username, err
}

const updateNote = `-- name: UpdateNote :one
UPDATE notes
SET
  title = COALESCE($2, title),
  text = COALESCE($3, text),
  updated_at = COALESCE($4, updated_at)
WHERE
  id = $1
RETURNING id
`

type UpdateNoteParams struct {
	ID        uuid.UUID
	Title     sql.NullString
	Text      sql.NullString
	UpdatedAt sql.NullTime
}

func (q *Queries) UpdateNote(ctx context.Context, arg *UpdateNoteParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, updateNote,
		arg.ID,
		arg.Title,
		arg.Text,
		arg.UpdatedAt,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}