package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Post Model, you can put it into a seperate
// package as well
type Post struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Comments  []Comment `json:"comments"`
}

type FeedPost struct {
	ID           int64    `json:"id"`
	Title        string   `json:"title"`
	Content      string   `json:"content"`
	UserID       int64    `json:"user_id"`
	Tags         []string `json:"tags"`
	CreatedAt    string   `json:"created_at"`
	Username     string   `json:"username"`
	CommentCount int      `json:"comment_count"`
}

type postsRepository interface {
	Create(context.Context, *Post) error
	GetByID(context.Context, int64) (*Post, error)
	Update(context.Context, *Post) error
	Delete(context.Context, *Post) error

	// users' personal feed
	Feed(context.Context, int64, *FeedPagination) ([]FeedPost, error)
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

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

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

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

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

func (ps *pqPosts) deletePost(ctx context.Context, tx *sql.Tx, post *Post) error {
	query := `
		DELETE FROM posts p WHERE p.id = $1;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := tx.ExecContext(
		ctx,
		query,
		post.ID,
	)
	if err != nil {
		return err
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if ra == 0 {
		return ErrNotFound
	}

	return nil
}

func (ps *pqPosts) deletePostComments(ctx context.Context, tx *sql.Tx, post *Post) error {
	query := `
		DELETE FROM comments c WHERE c.post_id = $1;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := tx.ExecContext(
		ctx,
		query,
		post.ID,
	)
	if err != nil {
		return err
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if ra == 0 {
		return ErrNotFound
	}

	return nil
}

func (ps *pqPosts) Delete(ctx context.Context, post *Post) error {
	return withTx(ps.db, ctx, func(tx *sql.Tx) error {
		// delete comments by post id
		if err := ps.deletePostComments(ctx, tx, post); err != nil {
			return err
		}

		// delete the post
		if err := ps.deletePost(ctx, tx, post); err != nil {
			return err
		}

		return nil
	})
}

func (ps *pqPosts) Update(ctx context.Context, post *Post) error {
	query := `
		UPDATE posts p
		SET
			title = $1,
			"content" = $2,
			tags = $3
		WHERE
			id = $4 AND
			user_id = $5;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := ps.db.ExecContext(
		ctx,
		query,
		post.Title,
		post.Content,
		pq.Array(post.Tags),
		post.ID,
		post.UserID,
	)
	if err != nil {
		return err
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if ra == 0 {
		return ErrNotFound
	}

	return nil
}

func (ps *pqPosts) Feed(ctx context.Context, userID int64, fp *FeedPagination) ([]FeedPost, error) {
	query := `
		SELECT
			p.id,
			p.title,
			p."content",
			p.tags,
			p.user_id,
			p.created_at,
			u.username,
			COUNT(c.id) AS comment_count
		FROM posts p
		JOIN users u ON u.id = p.user_id
		JOIN followers f ON f.user_id = p.user_id OR p.user_id = $1
		LEFT JOIN "comments" c ON c.post_id = p.id
		WHERE
			f.follower_id = $1 AND
			(
				p.title ILIKE '%' || $2 || '%' OR
				p.content ILIKE '%' || $2 || '%'
			) AND
			(
				ARRAY_LENGTH($3::text[], 1) IS NULL OR p.tags @> $3
			)
		GROUP BY p.id, u.username
		ORDER BY p.id ` + fp.Sort + `
		LIMIT $4 OFFSET $5;
		`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := ps.db.QueryContext(
		ctx,
		query,
		userID,
		fp.Search,
		pq.Array(fp.Tags),
		fp.Limit,
		fp.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fps := []FeedPost{}
	for rows.Next() {
		var fp FeedPost
		err = rows.Scan(
			&fp.ID,
			&fp.Title,
			&fp.Content,
			pq.Array(&fp.Tags),
			&fp.UserID,
			&fp.CreatedAt,
			&fp.Username,
			&fp.CommentCount,
		)
		if err != nil {
			return nil, err
		}

		fps = append(fps, fp)
	}

	return fps, nil
}
