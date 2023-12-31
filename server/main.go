package main

import (
	agents_handler "valorant-agents/handlers/agents"
	comments_handler "valorant-agents/handlers/comments"
	players_handler "valorant-agents/handlers/players"
	storage_handler "valorant-agents/handlers/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)


var PORT string = ":5000"

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins: "http://localhost:5173",
	}))
	app.Static("/storage/avatars", "./storage/avatars")

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON("Hello World")
	})
	//agents handlers
	agentsGroup := app.Group("/agents") 
	agentsGroup.Get("/", agents_handler.GetAgentsHandler)
	agentsGroup.Get("/:id", agents_handler.GetAgentHandler)
	agentsGroup.Post("/filter", agents_handler.FilterAgentsHandler)

	//comment handlers 
	commentGroup := app.Group("/comments")
	commentGroup.Get("/", comments_handler.GetCommentsHandler)
	commentGroup.Post("/create", comments_handler.CreateCommentHandler)
	commentGroup.Get("/:id/replies/", comments_handler.GetCommentRepliesHandler)
	//players handlers
	playersGroup := app.Group("/players")
	playersGroup.Get("/auth", players_handler.AuthPlayerHandler)
	playersGroup.Post("/register", players_handler.RegisterPlayerHandler)
	playersGroup.Post("/login", players_handler.LoginPlayerHandler)
	playersGroup.Get("/logout", players_handler.LogoutPlayerHandler)
	playersGroup.Get("/:username", players_handler.GetPlayerByUsernameHandler)
	playersGroup.Get("/generate/password", players_handler.GeneratePasswordHandler)
	//storage handlers
	storageGroup := app.Group("/storage")
	storageGroup.Post("/upload", storage_handler.UploadImageHandler)
	storageGroup.Delete("/delete/:image_name", storage_handler.DeleteImageHandler)
	
	
	app.Listen(PORT)
}
