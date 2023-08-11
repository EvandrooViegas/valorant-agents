package handlers

import (
	"encoding/json"
	"valorant-agents/utils"

	"github.com/gofiber/fiber/v2"
)



func GetAgentsHandler(c *fiber.Ctx) error {
	agents, err := GetAgents()
	utils.HandleErr(err)
	return c.JSON(agents)
}

func PostAgentsHandler(c *fiber.Ctx) error {
	body := c.Body()

	var request FilterRequest
	err := json.Unmarshal(body, &request)
	utils.HandleErr(err)
	filteredAgents := FilterAgents(request.Filter)
	return c.JSON(filteredAgents)
}
