package transaction_controller

import (
	"fmt"
	"time"
	"wsmail25/model"

	"github.com/gofiber/fiber/v2"
)

// func (u *TransactionHandler) InsertTransaction(ctx *fiber.Ctx) (err error) {
// 	var trans model.Transaction
// 	if err := ctx.BodyParser(&trans); err != nil {
// 		return err
// 	}

// 	err = u.transaction.InsertTransaction(trans)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (t *TransactionHandler) GetAllTransaction(ctx *fiber.Ctx) (err error) {
	trans, err := t.transaction.GetAllTransaction(ctx.Context())
	if err != nil {
		return fmt.Errorf("failed to get transaction data: %w", err)
	}
	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "Data transaksi berhasil diambil",
		"data":	trans,
	})
}

func (r *TransactionHandler) InsertTransaction(c *fiber.Ctx) error {
	var trans model.Transaction

	if err := c.BodyParser(&trans); err != nil {
		return fmt.Errorf("failed to parse request body: %w", err)
	}

	// Set waktu otomatis
	trans.CreatedAt = time.Now()
	trans.UpdatedAt = time.Now()

	// Simpan ke MongoDB melalui repository (repo akan generate connote otomatis)
	savedTrans, err := r.transaction.InsertTransaction(c.Context(), trans)
	if err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}

	return c.JSON(fiber.Map{
		"status":      "success",
		"message":     "Transaction successfully created",
		"transaction": savedTrans,
	})
}