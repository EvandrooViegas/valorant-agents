package players_handler

import "valorant-agents/services"

type CreatePlayerRequest struct {
	Player services.RegisterPlayerRequest `json:"player"`
}

type LoginPlayerRequest struct {
	Player services.LoginPlayerRequest `json:"player"`
}