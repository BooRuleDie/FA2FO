CREATE TABLE IF NOT EXISTS comments (
	id bigserial PRIMARY KEY,
	post_id bigint NOT NULL,
	user_id bigint NOT NULL,
	content TEXT NOT NULL,
	created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (post_id) REFERENCES posts(id)
);

CREATE INDEX IF NOT EXISTS idx_comments_post_id ON comments(post_id);
CREATE INDEX IF NOT EXISTS idx_comments_user_id ON comments(user_id);
