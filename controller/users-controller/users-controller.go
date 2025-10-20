package users_controller

import (
	"fmt"
	// "wsmail25/model"
	// "wsmail25/model"

	"github.com/gofiber/fiber/v2"
)

// func (u *UserHandler) InsertUser(ctx *fiber.Ctx) (err error) {
// 	var users model.Users
// 	if err := ctx.BodyParser(&users); err != nil {
// 		return err
// 	}

// 	err = u.user.InsertUser(users)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (u *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := u.user.GetAllUsers(c.Context())
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}
	return c.JSON(fiber.Map{"status": "succes", "users": users})
}

// func (u *UserHandler) InsertUser(c *fiber.Ctx) error {
// 	var user model.Users
// 	if err := c.BodyParser(&user); err != nil {
// 		return fmt.Errorf("failed to parse request body: %w", err)
// 	}
// 	password
// }