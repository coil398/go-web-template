package persistence

import (
	"api/model"
	"api/repository"

	"github.com/jmoiron/sqlx"
)

type UserStore struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) repository.User {
	return &UserStore{db}
}

func (u *UserStore) GetUser(id int) (*model.User, error) {
	return &model.User{
		Name: "Test",
	}, nil
}
