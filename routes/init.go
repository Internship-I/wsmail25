package routes

import (
	transRepo "wsmail25/repository/trx"
	usersRepo "wsmail25/repository/users"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(db *mongo.Client) {
	TransactionRepository = transRepo.NewTransaksiTable(db)
	UsersRepository = usersRepo.NewPenggunaTable(db)
}

func Router(app *fiber.App) (err error) {
	api := app.Group("/api")

	err = UserRoutes(api)
	if err != nil {
		return err
	}
	return
}
