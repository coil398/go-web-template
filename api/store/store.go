package store

import (
	"api/persistence"
	"api/repository"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	User repository.User
}

func NewStore(db *sqlx.DB) *Store {
	user := persistence.NewUserStore(db)

	return &Store{
		User: user,
	}
}
