package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	// InsertUser(ctx *fiber.Ctx) (err error)
	GetAllUsers(c *fiber.Ctx) error
	GetAllTransaction(ctx *fiber.Ctx) (err error)
	InsertTransaction(c *fiber.Ctx) error
	GetByConnote(c *fiber.Ctx) error
	GetByDeliveryStatus(c *fiber.Ctx) error	
	DeleteTransaction(c *fiber.Ctx) error
	UpdateDeliveryStatus(c *fiber.Ctx) error
}

// type TransactionController interface {
// 	// InsertTransaction(ctx *fiber.Ctx) (err error)
// 	GetAllTransaction(ctx *fiber.Ctx) (err error)
// 	InsertTransaction(c *fiber.Ctx) error
// 	GetByConnote(c *fiber.Ctx) error
// 	GetByDeliveryStatus(c *fiber.Ctx) error
// 	// UpdateDeliveryStatus(c *fiber.Ctx) error
// }
