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