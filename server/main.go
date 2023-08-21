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
	app.Static("/storage/avatars", "./storage/avatars")

	//agents handlers
	agentsGroup := app.Group("/agents")
	agentsGroup.Get("/", agents_handler.GetAgentsHandler)
	agentsGroup.Get("/:id", agents_handler.GetAgentHandler)
	agentsGroup.Post("/filter", agents_handler.FilterAgentsHandler)

	//players handlers
	playersGroup := app.Group("/players")
	playersGroup.Post("/register", players_handler.RegisterPlayer)
	//storage handlers
	storageGroup := app.Group("/storage")
	storageGroup.Post("/upload", storage_handler.UploadImageHandler)
	storageGroup.Delete("/delete/:image_name", storage_handler.DeleteImageHandler)
	
	
	app.Listen(PORT)
}
