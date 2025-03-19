CREATE TABLE IF NOT EXISTS followers (
    user_id BIGINT NOT NULL,
    follower_id BIGINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, follower_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (follower_id) REFERENCES users(id)
);

CREATE INDEX followers_user_id_idx ON followers(user_id);
CREATE INDEX followers_follower_id_idx ON followers(follower_id);