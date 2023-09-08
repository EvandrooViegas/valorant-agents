package services

import "github.com/golang-jwt/jwt/v4"

type Player struct {
	ID          string `json:"id" bson:"_id"`
	Avatar      string `json:"avatar"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Description string `json:"description"`
}

type RegisterPlayerRequest struct {
	Avatar      string `json:"avatar"`
	Username    string `json:"username"`
	Description string `json:"description"`
	Password    string `json:"password"`
}

type LoginPlayerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DatabaseNewPlayer struct {
	Username    string `json:"username" bson:"username,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
	Password    string `json:"password" bson:"password,omitempty"`
	Avatar      string `json:"avatar" bson:"avatar,omitempty"`
}
type IDPlayerTokenClaim struct {
	ID string `json:"id" bson:"id"`
}
type PlayerTokenClaims struct {
	jwt.RegisteredClaims
	Player IDPlayerTokenClaim `json:"player"`
}

// comments

type Comment struct {
	ID       string    `json:"id" bson:"_id"`
	AgentID  string    `json:"agent_id"`
	ParentID string    `json:"parent_id"`
	Text     string    `json:"text"`
	Author   Player    `json:"player"`
	Replies  []Comment `json:"replies"`
}

type RequestComment struct {
	AgentID  string `json:"agent_id"`
	ParentID string `json:"parent_id"`
	Text     string `json:"text"`
}

type DatabaseComment struct {
	AgentID  string `json:"agent_id" bson:"agent_id"`
	ParentID string `json:"parent_id" bson:"parent_id"`
	Text     string `json:"text" bson:"text"`
	Author   Player `json:"player" bson:"author"`
}

type CreateCommentRequest struct {
	Comment RequestComment `json:"comment"`
}
