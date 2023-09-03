package players_handler

import (
	"encoding/json"
	"fmt"
	"time"
	"valorant-agents/services"
	"valorant-agents/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterPlayerHandler(c *fiber.Ctx) error {
	body := c.Body()
	var request CreatePlayerRequest
	err := json.Unmarshal(body, &request)
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{Status: 500, Message: "Error Parsing the data", Error: err})
	}

	foundPlayers, err := services.FilterPlayerWithUsername(request.Player.Username)
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{
			Status:  fiber.StatusInternalServerError,
			Message: "A error occured",
		})
	}
	if len(foundPlayers) > 0 {
		return utils.WriteJSON(c, utils.WriteJSONpayload{
			Status:  fiber.StatusOK,
			Message: "A player with the same username was found, choose another one",
		})
	}

	nPlayer, err := services.CreatePlayer(request.Player)
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{
			Status:  fiber.StatusInternalServerError,
			Message: "Couldn't Create Player",
		})
	}

	token, err := services.CreatePlayerToken(services.IDPlayerTokenClaim{
		ID: nPlayer.ID,
	})

	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{
			Status:  fiber.StatusInternalServerError,
			Message: "Could not create token",
			Error:   err,
		})
	}
	fmt.Println("tokne created: ", token)
	c.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: token,
		Expires: time.Now().Add(24 * time.Hour),
		Path: "/",
	})
	return nil
}

func GeneratePasswordHandler(c *fiber.Ctx) error {
	password := GeneratePassword()
	return c.JSON(password)
}
