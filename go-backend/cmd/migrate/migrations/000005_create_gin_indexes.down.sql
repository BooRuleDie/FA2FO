DROP INDEX IF EXISTS idx_posts_tags;

DROP INDEX IF EXISTS idx_posts_title;
DROP INDEX IF EXISTS idx_comments_content;
DROP INDEX IF EXISTS idx_users_username;

DROP EXTENSION IF EXISTS pg_trgm;
