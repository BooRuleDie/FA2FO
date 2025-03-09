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