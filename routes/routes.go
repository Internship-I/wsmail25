package routes

import (
	transaction_controller "wsmail25/controller/transaction-controller"
	users_controller "wsmail25/controller/users-controller"

	"github.com/gofiber/fiber/v2"
)

func GetHome(ctx *fiber.Ctx) error {
	ipAddress := ctx.IP()
	if ipAddress == "" {
		ipAddress = "Unknown"
	}

	return ctx.JSON(fiber.Map{
		"ip_address": ipAddress,
	})
}

func UserRoutes(grp fiber.Router) (err error) {
	users := users_controller.NewUserController(UsersRepository)
	trans := transaction_controller.NewTransController(TransactionRepository)

	grp.Get("/", GetHome)
	groupes := grp.Group("/user")

	groupes.Get("/getallusers", users.GetAllUsers)
	groupes.Post("/inserttrans", trans.InsertTransaction)
	groupes.Get("/getalltransactions", trans.GetAllTransaction)
	groupes.Get("/getbyconnote/:connote", trans.GetByConnote)

	return
}

// func TransactionRoutes(grp fiber.Router) (err error) {
// 	trans := transaction_controller.NewTransController(TransactionRepository)

// 	grp.Get("/", GetHome)
// 	groupes := grp.Group("/transaction")

// 	groupes.Post("/inserttrans", trans.InsertTransaction)
// 	groupes.Get("/getalltransactions", trans.GetAllTransaction)
// 	return
// }

// func UserRoutes(grp fiber.Router) (err error) {
// 	users := users_controller.NewUserController(UsersRepository)
// 	trans := transaction_controller.NewTransController(TransactionRepository)

// 	userGroup := grp.Group("/user")
// 	transGroup := grp.Group("/transaction")

// 	userGroup.Post("/insert", users.InsertUser)
// 	userGroup.Get("/list", users.GetAllUsers)

// 	transGroup.Post("/insert", trans.InsertTransaction)
// 	transGroup.Get("/list", trans.GetAllTransaction)

// 	return
// }
