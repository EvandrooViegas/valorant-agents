package services

import "github.com/golang-jwt/jwt/v4"

type Player struct {
	ID          string `json:"_id" bson:"_id"`
	Avatar      string `json:"avatar"`
	Username    string `json:"username"`
	Password string `json:"password"`
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
	Password    string `json:"password"`
}

type DatabaseNewPlayer struct {
	Username    string `json:"username" bson:"username,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
	Password    string `json:"password" bson:"password,omitempty"`
	Avatar      string `json:"avatar" bson:"avatar,omitempty"`
}
type IDPlayerTokenClaim struct {
	ID string   `json:"id" bson:"id"`
}
type PlayerTokenClaims struct {
	jwt.RegisteredClaims
	Player IDPlayerTokenClaim `json:"player"`
}