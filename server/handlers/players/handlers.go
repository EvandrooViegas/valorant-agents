package players_handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func RegisterPlayer(c *fiber.Ctx) error {
	body := c.Body()
	fmt.Println(string(body))
	return fmt.Errorf("s")
}