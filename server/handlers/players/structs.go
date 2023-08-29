package players_handler

import "valorant-agents/services"

type CreatePlayerRequest struct {
	Player services.RequestPlayer `json:"player"`
}
