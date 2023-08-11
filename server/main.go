package main

import (

	handlers "valorant-agents/handlers/agents"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)


var PORT string = ":8080"

func main() {
	app := fiber.New()
	app.Use(cors.New())

	agentsGroup := app.Group("/agents")
	agentsGroup.Get("/", handlers.GetAgentsHandler)
	agentsGroup.Post("/", handlers.PostAgentsHandler)


	app.Listen(PORT)
}
