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
