package players_handler

import (
	"encoding/json"
	"valorant-agents/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterPlayerHandler(c *fiber.Ctx) error {
	body := c.Body()
	var request CreatePlayerRequest 
	err := json.Unmarshal(body, &request)
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{ Status: 500, Message: "Error Parsing the data", Error: err })
	}
	return utils.WriteJSON(c, utils.WriteJSONpayload{ Status: 200, Message: "Success", Data: body })
}

func GeneratePasswordHandler(c *fiber.Ctx) error {
	password := GeneratePassword()
	return c.JSON(password)
}