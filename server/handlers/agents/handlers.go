package agents_handler

import (
	"encoding/json"
	"valorant-agents/utils"
	"github.com/gofiber/fiber/v2"
)



func GetAgentHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	agent, error := GetAgent(id)
	utils.HandleErr(error)
	return c.JSON(agent)
}

func GetAgentsHandler(c *fiber.Ctx) error {
	agents, err := GetAgents()
	utils.HandleErr(err)
	return c.JSON(agents)
}

func FilterAgentsHandler(c *fiber.Ctx) error {
	body := c.Body()
	var request FilterRequest
	err := json.Unmarshal(body, &request)
	utils.HandleErr(err)
	filteredAgents := FilterAgents(request.Filter)
	return c.JSON(filteredAgents)
}
