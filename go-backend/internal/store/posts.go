package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

// Post Model, you can put it into a seperate
// package as well
type Post struct {
	ID        int64         `json:"id"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	UserID    int64         `json:"user_id"`
	Tags      []string      `json:"tags"`
	CreatedAt time.Duration `json:"created_at"`
	UpdatedAt time.Duration `json:"updated_at"`
}

type postsRepository interface {
	Create(context.Context, *Post) error
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

func (ps *pqPosts) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts(content, title, user_id, tags)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	err := ps.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	return err
}
