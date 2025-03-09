package store

import (
	"context"
	"database/sql"
)

// User Model, could be in a
// seperate package
type User struct {
	ID        int64          `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	// password field is excluded from JSON marshaling with `json:"-"`
	Password  string         `json:"-"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type usersRepository interface {
	Create(context.Context, *User) error
}

// postgreSQL Users struct that'll satisfy
// the usersRepository interface
type pqUsers struct {
	db *sql.DB
}

func newUsersRepo(db *sql.DB) *pqUsers {
	return &pqUsers{
		db: db,
	}
}

func (us *pqUsers) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO Users(username, password, email)
		VALUES($1, $2, $3) RETURNING id, created_at, updated_at
	`

	err := us.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		// should be hashed before inserting
		// or let database handle the hashing
		user.Password,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return err
}
