package players_handler

import (
	"encoding/json"
	"valorant-agents/services"
	"valorant-agents/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterPlayerHandler(c *fiber.Ctx) error {
	body := c.Body()
	var request CreatePlayerRequest
	err := json.Unmarshal(body, &request)
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{Status: fiber.StatusBadRequest, Message: "Error Parsing the data", Error: err})
	}

	player, err := services.RegisterPlayer(request.Player)
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{
			Status:  fiber.StatusBadRequest,
			Message: "Couldn't Create Player, ",
		})
	}
	if player.AlreadyExists {
		return utils.WriteJSON(c, utils.WriteJSONpayload{
			Status:  fiber.StatusAccepted,
			Message: "A player with the same username was found, choose another one",
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   player.Token,
		Expires: services.TOKEN_AND_COOKIE_EXPIRITY_TIME,
		Path:    "/",
	})
	responseData := map[string]string{
		"id": player.ID,
	}
	response := utils.WriteJSONpayload{
		Status:  fiber.StatusCreated,
		Message: "Player Created",
		Data:    responseData,
	}
	return utils.WriteJSON(c, response)
}

func LoginPlayerHandler(c *fiber.Ctx) error {
	// read the body
	body := c.Body()
	var request LoginPlayerRequest
	err := json.Unmarshal(body, &request)
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{
			Status:  fiber.StatusBadRequest,
			Message: "Error Parsing the data",
		})
	}
	// read the token and get the id from the cookies
	foundPlayer, err := services.LoginPlayer(request.Player)
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{
			Status:  fiber.StatusAccepted,
			Message: "Wrong Credentials",
		})
	}

	return utils.WriteJSON(c, utils.WriteJSONpayload{
		Status:  fiber.StatusOK,
		Message: "Login Successfully",
		Data: map[string]interface{}{
			"player": foundPlayer,
		},
	})
}

func AuthPlayerHandler(c *fiber.Ctx) error {
	token := c.Cookies("token")
	claims, err := services.ReadPlayerToken(token)
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{
			Status:  fiber.StatusUnauthorized,
			Message: "The token is invalid",
		})
	}
	id := claims.ID
	player, err := services.GetPlayerByID(id)
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{
			Status:  fiber.StatusInternalServerError,
			Message: "Could not fetch the user",
		})
	}
	return utils.WriteJSON(c, utils.WriteJSONpayload{
		Status:  fiber.StatusOK,
		Message: "Fetched the user successfully",
		Data: map[string]interface{}{
			"player": player,
		},
	})
}
func GeneratePasswordHandler(c *fiber.Ctx) error {
	password := GeneratePassword()
	return c.JSON(password)
}
