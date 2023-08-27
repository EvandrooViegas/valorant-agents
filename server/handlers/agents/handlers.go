package agents_handler

import (
	"encoding/json"
	"valorant-agents/utils"
	"github.com/gofiber/fiber/v2"
)



func GetAgentHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	agent, err := GetAgent(id)
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{ Status: 500, Message: "Error Getting the agent", Error: err })
	}
	return utils.WriteJSON(c, utils.WriteJSONpayload{ Status: 200, Message: "Success", Data: agent })
}

func GetAgentsHandler(c *fiber.Ctx) error {
	agents, err := GetAgents()
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{ Status: 500, Message: "Error Getting the agents", Error: err })
	}
	return utils.WriteJSON(c, utils.WriteJSONpayload{ Status: 200, Message: "Success", Data: agents })
}

func FilterAgentsHandler(c *fiber.Ctx) error {
	body := c.Body()
	var request FilterRequest
	err := json.Unmarshal(body, &request)
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{ Status: 500, Message: "Error Parsering the agents", Error: err })
	}
	filteredAgents, err := FilterAgents(request.Filter)
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{ Status: 500, Message: "Error Filtering the agents", Error: err })
	}
	return utils.WriteJSON(c, utils.WriteJSONpayload{ Status: 200, Message: "Success", Data: filteredAgents })
}
