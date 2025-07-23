package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("not found")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Users    usersRepository
	Posts    postsRepository
	Comments commentsRepository
	Roles    rolesRepository
}

func NewPostgreSQLStorage(db *sql.DB) Storage {
	users := newUsersRepo(db)
	posts := newPostsRepo(db)
	comments := newCommentsRepo(db)
	roles := newRolesRepo(db)

	return Storage{
		Users:    users,
		Posts:    posts,
		Comments: comments,
		Roles:    roles,
	}
}

func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return errors.New("transaction error: " + err.Error() + ", rollback error: " + rbErr.Error())
		}
		return err
	}

	return tx.Commit()
}
