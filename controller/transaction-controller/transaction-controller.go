package transaction_controller

import (
	"fmt"
	"time"
	"wsmail25/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// ✅ Fungsi Update Status Pengiriman (umum)
func (h *TransactionHandler) UpdateDeliveryStatus(c *fiber.Ctx) error {
	connote := c.Params("connote")

	var req struct {
		Status string `json:"status"`
	}

	// Parsing JSON dari request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal membaca request body",
			"error":   err.Error(),
		})
	}

	if req.Status == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Field 'status' wajib diisi",
		})
	}

	// Check that transaction exists
	trx, err := h.transaction.GetByConnote(c.Context(), connote)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Transaksi dengan connote '%s' tidak ditemukan: %v", connote, err),
		})
	}

	// NOTE: repository does not expose an UpdateDeliveryStatus method.
	// Return Not Implemented and include the existing transaction so caller can implement persistence.
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"status":  "error",
		"message": "Repository belum mengimplementasikan UpdateDeliveryStatus; tambahkan method pada repository untuk menyimpan perubahan status",
		"data":    trx,
	})
}

// ✅ Fungsi Kirim WA ketika "On Delivery"
func (h *TransactionHandler) SendWAOnDelivery(c *fiber.Ctx) error {
	var payload struct {
		ID string `json:"id"`
	}

	// Parsing payload JSON
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Payload tidak valid",
			"error":   err.Error(),
		})
	}

	// Validasi ObjectID
	_, err := primitive.ObjectIDFromHex(payload.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "ID tidak valid",
			"error":   err.Error(),
		})
	}

	// Repository belum mengimplementasikan fungsi untuk mengirim WA; kembalikan Not Implemented agar pemanggil mengetahui
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"status":  "error",
		"message": "Repository belum mengimplementasikan metode untuk mengirim WA ('SendWADelivered' atau 'SendWAOnDelivery'); tambahkan metode pada repository untuk melakukan pengiriman dan update status",
		"payload": map[string]interface{}{"id": payload.ID},
	})

	// return c.JSON(fiber.Map{
	// 	"status":  "success",
	// 	"message": fmt.Sprintf("WA 'On Delivery' berhasil dikirim untuk %s", updated.ConsignmentNote),
	// 	"data":    updated,
	// })
}

// ✅ Fungsi Kirim WA ketika "Delivered"
func (h *TransactionHandler) SendWADelivered(c *fiber.Ctx) error {
	var payload struct {
		ID string `json:"id"`
	}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Payload tidak valid",
			"error":   err.Error(),
		})
	}

	_, err := primitive.ObjectIDFromHex(payload.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "ID tidak valid",
			"error":   err.Error(),
		})
	}

	// Repository belum mengimplementasikan fungsi untuk mengirim WA; kembalikan Not Implemented agar pemanggil mengetahui
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"status":  "error",
		"message": "Repository belum mengimplementasikan metode SendWADelivered; tambahkan metode pada repository untuk melakukan pengiriman WA dan update status",
		"payload": map[string]interface{}{"id": payload.ID},
	})
}