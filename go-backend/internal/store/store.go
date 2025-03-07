package store

import (
	"database/sql"
)

type Storage struct {
	Users usersRepository
	Posts postsRepository
}

func NewPostgreSQLStorage(db *sql.DB) Storage {
	users := newUsersRepo(db)
	posts := newPostsRepo(db)

	return Storage{
		Users: users,
		Posts: posts,
	}
}
