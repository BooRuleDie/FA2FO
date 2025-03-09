# Go Migrate Installation

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

# Notes

In PostgreSQL, using the TEXT data type for fields is not an anti-pattern; instead, it's the recommended option. In MySQL, it can cause performance issues, but in PostgreSQL, it can be thought of as an elastic VARCHAR (TEXT is actually similar to VARCHAR but without length limitation). As long as you don't need any text length restriction, there is no need for you to use VARCHAR.