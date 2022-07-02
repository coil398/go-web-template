package handler

import (
	"api/store"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	store *store.Store
}

func NewHandlers(store *store.Store) *Handlers {
	return &Handlers{store}
}

func (h Handlers) PostUser(ctx echo.Context) error {
	return nil
}

func (h *Handlers) GetUsersUserId(ctx echo.Context, userId int) error {
	return nil
}

func (h *Handlers) PatchUsersUserId(ctx echo.Context, userId int) error {
	return nil
}
