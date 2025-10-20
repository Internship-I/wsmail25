package users_controller

import (
	// "wsmail25/controller"
	"wsmail25/repository"
)

type UserHandler struct {
	user repository.UsersRepository
}

func NewUserController(user repository.UsersRepository) *UserHandler {
	return &UserHandler{
		user: user,
	}
}
