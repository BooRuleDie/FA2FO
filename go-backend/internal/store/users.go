package store

import (
	"context"
	"database/sql"
)

type usersRepository interface {
	Create(context.Context) error
}

// postgreSQL Users struct that'll satisfy
// the usersRepository interface
type pqUsers struct {
	db *sql.DB
}

func newUsersRepo(db *sql.DB) *pqUsers {
	return &pqUsers{
		db: db,
	}
}

func (ps *pqUsers) Create(ctx context.Context) error {
	
	return nil
}
