# Go Migrate Installation

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

# Notes

## PostgreSQL Text Field
In PostgreSQL, using the TEXT data type for fields is not an anti-pattern; instead, it's the recommended option. In MySQL, it can cause performance issues, but in PostgreSQL, it can be thought of as an elastic VARCHAR (TEXT is actually similar to VARCHAR but without length limitation). As long as you don't need any text length restriction, there is no need for you to use VARCHAR.

## PostgreSQL Single vs Double Quotes

Unlike MySQL, PostgreSQL has distinct uses for single quotes (`'`) and double quotes (`"`):

```sql
-- Correct: Using single quotes for string values
SELECT * FROM users WHERE name = 'John';
INSERT INTO products (name, description) VALUES ('Laptop', 'A portable computer');

-- Escaping single quotes within strings
SELECT * FROM comments WHERE text = 'John''s comment';
```

```sql
-- Using double quotes for table/column names with special cases
SELECT "firstName" FROM "User Accounts";

-- Double quotes required for case-sensitivity or reserved keywords
CREATE TABLE "Order" ("itemId" integer, "createdAt" timestamp);
```

Using the wrong quote style will either cause syntax errors or produce unexpected results in your queries.

### PostgreSQL & MySQL Index on Foreign Keys

Both PostgreSQL and MySQL create an index for the primary key on tables. However, when it comes to foreign keys, there is a difference. MySQL creates the index automatically just like it does on primary keys, however PostgreSQL expects you to create the index manually for the foreign key. Here's the recommended approach in PostgreSQL:

```sql
CREATE TABLE users (
    id INT PRIMARY KEY,
    name VARCHAR(100)
);

CREATE TABLE orders (
    order_id INT PRIMARY KEY,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id),
    -- Other columns...
);

-- Create index on foreign key for performance
CREATE INDEX idx_orders_user_id ON orders(user_id);
```

### Close on sql.Rows in Go

When using `QueryRow` or `QueryRowContext`, you don't need to worry about closing the result as it's handled automatically. However, when using `Query` or `QueryContext` which return `*sql.Rows`, you must explicitly close the rows using `defer rows.Close()`. Here's an example:

```go
// Using Query - requires explicit Close()
rows, err := db.Query("SELECT id, name FROM users")
if err != nil {
    return err
}
defer rows.Close() // Important!

for rows.Next() {
    // Process rows
}

// Using QueryRow - no Close() needed
row := db.QueryRow("SELECT id, name FROM users WHERE id = $1", id)
var user User
err = row.Scan(&user.ID, &user.Name)
```

### Group By Behavior: PostgreSQL vs MySQL

PostgreSQL strictly follows the SQL standard for GROUP BY operations. When using GROUP BY, you must:
* Include ALL non-aggregated columns that appear in the SELECT clause in the GROUP BY clause
* Or use aggregate functions (COUNT, SUM, json_agg, etc.) for columns not in GROUP BY

```sql
-- PostgreSQL: This will NOT work
SELECT
    p.id,
    p.title,      -- Error: must be in GROUP BY
    p.content,    -- Error: must be in GROUP BY
    json_agg(c.*) as comments
FROM posts p
LEFT JOIN comments c ON c.post_id = p.id
GROUP BY p.id;

-- PostgreSQL: This will work
SELECT
    p.id,
    p.title,
    p.content,
    json_agg(c.*) as comments
FROM posts p
LEFT JOIN comments c ON c.post_id = p.id
GROUP BY p.id, p.title, p.content;
```

MySQL is more lenient with GROUP BY operations by default. It will:
* Allow columns in SELECT that are not in GROUP BY
* Automatically select a value (usually the first one it finds) for non-grouped columns

```sql
-- MySQL: This will work
SELECT
    p.id,
    p.title,      -- MySQL picks a value
    p.content,    -- MySQL picks a value
    JSON_ARRAYAGG(
        JSON_OBJECT(
            'id', c.id,
            'content', c.content
        )
    ) as comments
FROM posts p
LEFT JOIN comments c ON c.post_id = p.id
GROUP BY p.id;
```

### PostgreSQL Index Types

| Field Type | Best Index Type | Common Query Patterns |
|------------|----------------|---------------------|
| Integer (e.g., id, price, age) | B-Tree | `=`, `<`, `>`, `BETWEEN`, `ORDER BY` |
| Text (e.g., name, description) | B-Tree (exact match)<br>GiST (trigram) | `=`, `LIKE 'abc%'`<br>`ILIKE '%abc%'` |
| Full-Text Search (e.g., bio, content) | GIN (tsvector) | `@@`, `to_tsquery()` |
| Arrays (e.g., text[]) | GIN | `@>`, `&&`, `<@` |
| JSONB | GIN (jsonb)<br>GIN (jsonb_path_ops) | `@>`, `?`, `?&`, `?|`<br>JSON path operators |
