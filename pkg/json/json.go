package json

import "github.com/gofiber/fiber/v2"

type ReturnData struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

func (respdata ReturnData) WriteToBody(ctx *fiber.Ctx) error {
	return ctx.Status(respdata.Code).JSON(respdata)
}
