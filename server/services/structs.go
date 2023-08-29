package services
type Player struct {
	ID          string  `db:"id" json:"id"` 
	Avatar      string `json:"avatar"`
	Username    string `json:"username"`
	Description string `json:"description"`
}

type RequestPlayer struct {
	Avatar      string `json:"avatar"`
	Username    string `json:"username"`
	Description string `json:"description"`
	Password    string `json:"password"`
}

type DatabaseNewPlayer struct {
	Username    string `json:"username" bson:"username,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
	Password    string `json:"password" bson:"password,omitempty"`
	Avatar      string `json:"avatar" bson:"avatar,omitempty"`
}
