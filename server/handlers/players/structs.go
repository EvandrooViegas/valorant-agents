package players_handler 

type RequestPlayer struct {
	Avatar string `json:"avatar"`
	Username string `json:"username"`
	Description string `json:"description"`
	Password string `json:"password"`
}

type CreatePlayerRequest struct {
	Player RequestPlayer `json:"player"`
}