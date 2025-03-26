package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidUserID     = errors.New("invalid userID")
	ErrAlreadyFollowing  = errors.New("already following this user")
	ErrDuplicateEmail    = errors.New("email address already exists")
	ErrDuplicateUsername = errors.New("username already exists")
	ErrTokenExpired      = errors.New("token expired")
)

// User Model, could be in a
// seperate package
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// password field is excluded from JSON marshaling with `json:"-"`
	Password  password `json:"-"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type password struct {
	text *string
	hash []byte
}

func (p *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.text = &text
	p.hash = hash

	return nil
}

type usersRepository interface {
	Create(context.Context, *sql.Tx, *User) error
	GetByID(context.Context, int64) (*User, error)

	// userID, followerID
	Follow(context.Context, int64, int64) error
	Unfollow(context.Context, int64, int64) error

	// auth
	CreateAndInvite(context.Context, *User, string, time.Duration) error
	Activate(context.Context, string) error
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

func (us *pqUsers) Create(ctx context.Context, tx *sql.Tx, user *User) error {
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
		user.Password.hash,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		case err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"`:
			return ErrDuplicateUsername
		default:
			return err
		}
	}

	return nil
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

func (us *pqUsers) Follow(ctx context.Context, userID int64, followerID int64) error {
	query := `
		INSERT INTO followers (user_id, follower_id) VALUES ($1, $2);
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := us.db.ExecContext(
		ctx,
		query,
		userID,
		followerID,
	)

	// handle follow conflicts
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrAlreadyFollowing
		}
		return err
	}

	return nil
}

func (us *pqUsers) Unfollow(ctx context.Context, userID int64, followerID int64) error {
	query := `
		DELETE FROM followers WHERE user_id = $1 AND follower_id = $2;
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := us.db.ExecContext(
		ctx,
		query,
		userID,
		followerID,
	)

	return err
}

func (us *pqUsers) CreateAndInvite(ctx context.Context, user *User, token string, invitationExp time.Duration) error {
	return withTx(us.db, ctx, func(tx *sql.Tx) error {
		// create the user
		if err := us.Create(ctx, tx, user); err != nil {
			return err
		}

		// create the invitation
		if err := us.createUserInvitation(ctx, tx, token, invitationExp, user.ID); err != nil {
			return err
		}

		return nil
	})
}

func (us *pqUsers) createUserInvitation(ctx context.Context, tx *sql.Tx, token string, invitationExp time.Duration, userID int64) error {
	query := `
		INSERT INTO user_invitations (token, user_id, expiry) VALUES($1, $2, $3);
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, token, userID, time.Now().Add(invitationExp))
	if err != nil {
		return err
	}

	return nil
}

func (us *pqUsers) findUserByToken(ctx context.Context, tx *sql.Tx, token string) (int64, error) {
	query := `
		SELECT
		    ui.user_id,
		    CASE
		        WHEN NOW() > ui.expiry THEN TRUE
		        ELSE FALSE
		    END AS is_expired
		FROM user_invitations ui
		WHERE ui.token = $1;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var userID int64
	var isExpired bool
	if err := tx.QueryRowContext(ctx, query, token).Scan(&userID, &isExpired); err != nil {
		if err == sql.ErrNoRows {
			return 0, ErrNotFound
		} else {
			return 0, err
		}
	}
	if isExpired == true {
		return 0, ErrTokenExpired
	}

	return userID, nil
}

func (us *pqUsers) updateUserActivation(ctx context.Context, tx *sql.Tx, userID int64) error {
	query := `
		UPDATE users SET is_active = TRUE WHERE id = $1;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, userID)

	return err
}

func (us *pqUsers) deleteInvitationToken(ctx context.Context, tx *sql.Tx, token string) error {
	query := `
		DELETE FROM user_invitations WHERE "token" = $1;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, token)

	return err
}

func (us *pqUsers) Activate(ctx context.Context, token string) error {
	return withTx(us.db, ctx, func(tx *sql.Tx) error {
		// find users who owns the token
		userID, err := us.findUserByToken(ctx, tx, token)
		if err != nil {
			return err
		}

		// update the user
		err = us.updateUserActivation(ctx, tx, userID)
		if err != nil {
			return err
		}

		// clean the invitation token
		err = us.deleteInvitationToken(ctx, tx, token)
		if err != nil {
			return err
		}

		return nil
	})
}
