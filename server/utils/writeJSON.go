package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type WriteJSONpayload struct {
	Status  int
	Message string
	Data    any
	Error   any
}

func WriteJSON(c *fiber.Ctx, payload WriteJSONpayload) error {
	c.Status(payload.Status)
	if payload.Error != nil {
		fmt.Println("Error: ", payload.Error)

	}
	return c.JSON(fiber.Map{"status": payload.Status, "message": payload.Message, "data": payload.Data, "error": payload.Error})
}
