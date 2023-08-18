package main

import (
	agents_handler "valorant-agents/handlers/agents"
	players_handler "valorant-agents/handlers/players"
	storage_handler "valorant-agents/handlers/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)


var PORT string = ":8080"

func main() {
	app := fiber.New()
	app.Use(cors.New())

	agentsGroup := app.Group("/agents")
	agentsGroup.Get("/", agents_handler.GetAgentsHandler)
	agentsGroup.Get("/:id", agents_handler.GetAgentHandler)
	agentsGroup.Post("/filter", agents_handler.FilterAgentsHandler)

	playersGroup := app.Group("/players")
	playersGroup.Post("/register", players_handler.RegisterPlayer)

	storageGroup := app.Group("/storage")
	storageGroup.Post("/upload", storage_handler.UploadImage)
	app.Static("/storage/avatars", "./storage/avatars", fiber.Static{
		Compress: true,
		Download: true,
	})
	
	app.Listen(PORT)
}
