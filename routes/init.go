package routes

import (
	transRepo "wsmail25/repository/trans"
	// usersRepo "wsmail25/repository/users"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(db *mongo.Client) {
	// UsersRepository = usersRepo.NewUsersTable(db)
	TransactionRepository = transRepo.NewTransaksiTable(db)
}

func Router(app *fiber.App) (err error) {
	api := app.Group("/api") //semua route dibawah ini punya prefix /api

	err = UserRoutes(api)
	if err != nil {
		return err
	}
	return
}
