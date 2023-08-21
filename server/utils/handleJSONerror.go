package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func HandleJSONerror(c *fiber.Ctx, json fiber.Map) error {
	fmt.Println("Error: ", json)
	return c.JSON(json)
	
}