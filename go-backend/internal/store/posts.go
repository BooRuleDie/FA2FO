package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

var ErrNotFound = errors.New("not found")

// Post Model, you can put it into a seperate
// package as well
type Post struct {
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	UserID    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type postsRepository interface {
	Create(context.Context, *Post) error
	GetByID(context.Context, int64) (*Post, error)
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

func (ps *pqPosts) GetByID(ctx context.Context, postID int64) (*Post, error) {
	query := `
		SELECT
			p.id,
			p.title,
			p.content,
			p.user_id,
			p.tags,
			p.created_at,
			p.updated_at
		FROM posts p
		WHERE p.id = $1;
	`
	var post Post
	err := ps.db.QueryRowContext(
		ctx,
		query,
		postID,
	).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.UserID,
		pq.Array(&post.Tags),
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}
