package transaction_controller

import (
	"wsmail25/model"

	"github.com/gofiber/fiber/v2"
)

func (u *TransactionHandler) InsertTransaction(ctx *fiber.Ctx) (err error) {
	var trans model.Transaction
	if err := ctx.BodyParser(&trans); err != nil {
		return err
	}

	err = u.transaction.InsertTransaction(trans)
	if err != nil {
		return err
	}
	return nil
}

func (u *TransactionHandler) GetAllTransaction(ctx *fiber.Ctx) (err error) {
	var trans model.Transaction
	trans, err = u.transaction.GetAllTransaction()
	if err != nil {
		return err
	}
	return ctx.JSON(trans)
}