package store

import (
	"context"
	"database/sql"
)

type postsRepository interface {
	Create(context.Context) error
}

// postgreSQL Posts struct that'll satisfy
// the postsRepository interface
type pqPosts struct {
	db *sql.DB
}

func newPostsRepo(db *sql.DB) *pqPosts {
	return &pqPosts{
		db: db,
	}
}

func (ps *pqPosts) Create(ctx context.Context) error {

	return nil
}
