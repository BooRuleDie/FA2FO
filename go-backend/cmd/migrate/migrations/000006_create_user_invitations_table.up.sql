CREATE TABLE IF NOT EXISTS user_invitations (
    token BYTEA PRIMARY KEY,
    user_id BIGINT REFERENCES users(id) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_tokens_user_id ON user_invitations(user_id);