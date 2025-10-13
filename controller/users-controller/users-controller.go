package users_controller

import (
	"wsmail25/model"

	"github.com/gofiber/fiber/v2"
)

func (u *UserHandler) InsertUser(ctx *fiber.Ctx) (err error) {
	var users model.User
	if err := ctx.BodyParser(&users); err != nil {
		return err
	}

	err = u.user.InsertUser(users)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserHandler) GetAllUsers(ctx *fiber.Ctx) (err error) {
	var users model.User
	users, err = u.user.GetAllUsers()
	if err != nil {
		return err
	}
	return ctx.JSON(users)
}