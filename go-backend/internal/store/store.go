package store

import (
	"database/sql"
)

type Storage struct {
	Users usersRepository
	Posts postsRepository
	Comments commentsRepository
}

func NewPostgreSQLStorage(db *sql.DB) Storage {
	users := newUsersRepo(db)
	posts := newPostsRepo(db)
	comments := newCommentsRepo(db)

	return Storage{
		Users: users,
		Posts: posts,
		Comments: comments,
	}
}
