package repository

import "api/model"

type User interface {
	GetUser(id int) (*model.User, error)
}
