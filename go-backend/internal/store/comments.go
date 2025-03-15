package store

import (
	"context"
	"database/sql"
)

type Comment struct {
	ID        int64  `json:"id"`
	PostID    int64  `json:"post_id"`
	UserID    int64  `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
}

type commentsRepository interface {
	GetByPostID(context.Context, int64) ([]Comment, error)
	Create(context.Context, *Comment) error
}

type pqComments struct {
	db *sql.DB
}

func newCommentsRepo(db *sql.DB) *pqComments {
	return &pqComments{
		db: db,
	}
}

func (cs *pqComments) GetByPostID(ctx context.Context, postID int64) ([]Comment, error) {
	query := `
		SELECT
			c.id,
			c.post_id,
			c.user_id,
			c.content,
			c.created_at,
			u.id,
			u.username,
			u.email,
			u.created_at,
			u.updated_at
	FROM comments c
	JOIN users u ON u.id = c.user_id
	WHERE c.post_id = $1;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := cs.db.QueryContext(
		ctx,
		query,
		postID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []Comment{}
	for rows.Next() {
		var c Comment
		c.User = User{}

		err = rows.Scan(
			&c.ID,
			&c.PostID,
			&c.UserID,
			&c.Content,
			&c.CreatedAt,
			&c.User.ID,
			&c.User.Username,
			&c.User.Email,
			&c.User.CreatedAt,
			&c.User.UpdatedAt,
		)

		comments = append(comments, c)
	}

	return comments, nil
}

func (cs *pqComments) Create(ctx context.Context, comment *Comment) error {
	query := `
		INSERT INTO "comments"(post_id, user_id, "content") VALUES($1, $2, $3);
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := cs.db.ExecContext(
		ctx,
		query,
		comment.PostID,
		comment.UserID,
		comment.Content,
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
