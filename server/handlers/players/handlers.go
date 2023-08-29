package players_handler

import (
	"encoding/json"
	"fmt"
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
		return  utils.WriteJSON(c, utils.WriteJSONpayload{
			Status: fiber.StatusInternalServerError,
			Message: "A error occured",
		})
	}
	fmt.Println(foundPlayers)
	if len(foundPlayers) > 0 {
		return  utils.WriteJSON(c, utils.WriteJSONpayload{
			Status: fiber.StatusOK,
			Message: "A player with the same username was found, choose another one",
		})
	}

	nPlayer, err := services.CreatePlayer(request.Player)
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{
			Status: fiber.StatusInternalServerError,	
			Message: "Couldn't Create Player",
		})
	}	

	// should return token
	return utils.WriteJSON(c, utils.WriteJSONpayload{
		Status:  fiber.StatusCreated,
		Message: "Created Player Successfully",
		Data:   nPlayer,
	})


}

func GeneratePasswordHandler(c *fiber.Ctx) error {
	password := GeneratePassword()
	return c.JSON(password)
}
