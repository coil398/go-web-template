package usecase

import (
	"api/model"
	"api/store"
)

func GetUsersUserId(store *store.Store, id int) (*model.User, error) {
	user, err := store.User.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
