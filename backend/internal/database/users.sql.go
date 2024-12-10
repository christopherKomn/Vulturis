// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email,hashed_password,name,surname,phone,mobile,address)
VALUES (
    gen_random_uuid (),
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING id, created_at, updated_at, email, hashed_password, name, surname, phone, mobile, address
`

type CreateUserParams struct {
	Email          string
	HashedPassword string
	Name           string
	Surname        string
	Phone          string
	Mobile         string
	Address        string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Email,
		arg.HashedPassword,
		arg.Name,
		arg.Surname,
		arg.Phone,
		arg.Mobile,
		arg.Address,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.HashedPassword,
		&i.Name,
		&i.Surname,
		&i.Phone,
		&i.Mobile,
		&i.Address,
	)
	return i, err
}

const deleteUserByEmail = `-- name: DeleteUserByEmail :exec
delete from users WHERE email = $1
`

func (q *Queries) DeleteUserByEmail(ctx context.Context, email string) error {
	_, err := q.db.ExecContext(ctx, deleteUserByEmail, email)
	return err
}

const deleteUserById = `-- name: DeleteUserById :exec
delete from users WHERE id = $1
`

func (q *Queries) DeleteUserById(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserById, id)
	return err
}

const deleteUsers = `-- name: DeleteUsers :exec
delete  from users
`

func (q *Queries) DeleteUsers(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteUsers)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, created_at, updated_at, email, hashed_password, name, surname, phone, mobile, address FROM users WHERE $1 = users.email
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.HashedPassword,
		&i.Name,
		&i.Surname,
		&i.Phone,
		&i.Mobile,
		&i.Address,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, created_at, updated_at, email, hashed_password, name, surname, phone, mobile, address FROM users WHERE $1 = users.id
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.HashedPassword,
		&i.Name,
		&i.Surname,
		&i.Phone,
		&i.Mobile,
		&i.Address,
	)
	return i, err
}

const updateUserEmailById = `-- name: UpdateUserEmailById :one
UPDATE users SET email = $1
WHERE id = $2
RETURNING id, created_at, updated_at, email, hashed_password, name, surname, phone, mobile, address
`

type UpdateUserEmailByIdParams struct {
	Email string
	ID    uuid.UUID
}

func (q *Queries) UpdateUserEmailById(ctx context.Context, arg UpdateUserEmailByIdParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserEmailById, arg.Email, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.HashedPassword,
		&i.Name,
		&i.Surname,
		&i.Phone,
		&i.Mobile,
		&i.Address,
	)
	return i, err
}

const updateUserPasswordByEmail = `-- name: UpdateUserPasswordByEmail :one
UPDATE users SET hashed_password = $1
WHERE email = $2
RETURNING id, created_at, updated_at, email, hashed_password, name, surname, phone, mobile, address
`

type UpdateUserPasswordByEmailParams struct {
	HashedPassword string
	Email          string
}

func (q *Queries) UpdateUserPasswordByEmail(ctx context.Context, arg UpdateUserPasswordByEmailParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserPasswordByEmail, arg.HashedPassword, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.HashedPassword,
		&i.Name,
		&i.Surname,
		&i.Phone,
		&i.Mobile,
		&i.Address,
	)
	return i, err
}

const updateUsers = `-- name: UpdateUsers :exec
UPDATE users
SET email = $1, hashed_password = $2
WHERE id = $3
`

type UpdateUsersParams struct {
	Email          string
	HashedPassword string
	ID             uuid.UUID
}

func (q *Queries) UpdateUsers(ctx context.Context, arg UpdateUsersParams) error {
	_, err := q.db.ExecContext(ctx, updateUsers, arg.Email, arg.HashedPassword, arg.ID)
	return err
}