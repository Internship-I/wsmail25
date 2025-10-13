package exception

import (
	"wsmail25/pkg/json"

	"github.com/gofiber/fiber/v2"
)

func ErrHandler(ctx *fiber.Ctx, err error) error {
	resp := json.ReturnData{
		fiber.StatusInternalServerError,
		false,
		"Error terjadi" + err.Error(),
		nil,
	}
	if e, ok := err.(*fiber.Error); ok {
		resp.Code = e.Code
		resp.Status = e.Message
	}
	return resp.WriteToBody(ctx)
}
