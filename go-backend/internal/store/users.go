package store

import (
	"context"
	"database/sql"
	"errors"
)

var ErrInvalidUserID = errors.New("invalid userID")

// User Model, could be in a
// seperate package
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// password field is excluded from JSON marshaling with `json:"-"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type usersRepository interface {
	Create(context.Context, *User) error
	GetByID(context.Context, int64) (*User, error)
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

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

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

func (us *pqUsers) GetByID(ctx context.Context, userID int64) (*User, error) {
	query := `
		SELECT
			u.id,
			u.username,
			u.email,
			u.created_at,
			u.updated_at
		FROM users u
		WHERE u.id = $1;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var user User
	err := us.db.QueryRowContext(
		ctx,
		query,
		userID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}
