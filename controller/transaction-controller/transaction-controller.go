package transaction_controller

import (
	"fmt"
	"log"
	"time"
	"wsmail25/model"

	"github.com/gofiber/fiber/v2"
)

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

func (h *TransactionHandler) GetByConnote(c *fiber.Ctx) error {
    connote := c.Params("connote")
    trx, err := h.transaction.GetByConnote(c.Context(), connote)
    if err != nil {
        return c.Status(404).JSON(fiber.Map{"status":"error","message":err.Error()})
    }
    return c.JSON(fiber.Map{"status":"success","data":trx})
}

func (h *TransactionHandler) GetByDeliveryStatus(c *fiber.Ctx) error {
	status := c.Params("status")
	trxs, err := h.transaction.GetByDeliveryStatus(c.Context(), status)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status":"error","message":err.Error()})
	}
	return c.JSON(fiber.Map{"status":"success","data":trxs})
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

// ✅ Fungsi untuk update delivery status
func (c *TransactionHandler) UpdateDeliveryStatus(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Struktur data yang diterima dari FE
	var req struct {
		Status string `json:"status"`
		Reason string `json:"reason,omitempty"`
	}

	// Parsing body JSON dari FE
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "request body tidak valid: " + err.Error(),
		})
	}

	// Validasi status pengiriman
	validStatuses := map[string]bool{
		"On Delivery": true,
		"Delivered":   true,
		"Failed":      true,
	}

	if !validStatuses[req.Status] {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "status tidak valid, gunakan: On Delivery, Delivered, atau Failed",
		})
	}

	// Jalankan fungsi update di repository
		err := c.transaction.UpdateDeliveryStatus(ctx.Context(), id, req.Status, req.Reason)
		if err != nil {
			log.Println("[ERROR] gagal update status pengiriman:", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	// Kembalikan response sukses
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":        "Status pengiriman berhasil diperbarui",
		"delivery_id":    id,
		"new_status":     req.Status,
		"failure_reason": req.Reason,
		"updated_at":     time.Now(),
	})
}

// ✅ Delete Transaction by ID
func (r *TransactionHandler) DeleteTransaction(c *fiber.Ctx) error {
	id := c.Params("id")

	deleted, err := r.transaction.DeleteTransaction(c.Context(), id)
	if err != nil {
		log.Printf("[ERROR] gagal menghapus transaksi: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Transaksi dengan ID %s tidak ditemukan", id),
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("Transaksi dengan ID %s berhasil dihapus", id),
		"data":    deleted,
	})
}