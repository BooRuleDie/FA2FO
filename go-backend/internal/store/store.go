package store

import (
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
}

func NewPostgreSQLStorage(db *sql.DB) Storage {
	users := newUsersRepo(db)
	posts := newPostsRepo(db)
	comments := newCommentsRepo(db)

	return Storage{
		Users:    users,
		Posts:    posts,
		Comments: comments,
	}
}
