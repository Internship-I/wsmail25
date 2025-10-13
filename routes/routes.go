package routes

import (
	users_controller "wsmail25/controller/users-controller"
	transaction_controller "wsmail25/controller/transaction-controller"

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
	// matkul
	groupes.Post("/insertUs", users.InsertUser)
	groupes.Get("/User", users.GetAllUsers)
	groupes.Post("/insertTrans", trans.InsertTransaction)
	groupes.Get("/transaction", trans.GetAllTransaction)

	return
}
