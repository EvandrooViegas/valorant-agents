package handlers

type Role struct {
	Name string `json:"displayName"`
	Icon string `json:"displayIcon"`
}
type Abilities struct {
	Name        string `json:"displayName"`
	Description string `json:"description"`
	Icon        string `json:"displayIcon"`
}
type Agent struct {
	UUID        string      `json:"uuid"`
	DisplayName string      `json:"displayName"`
	Description string      `json:"description"`
	DisplayIcon string      `json:"displayIcon"`
	Portait     string      `json:"fullPortrait"`
	Background  string      `json:"background"`
	Role        Role        `json:"role"`
	Abilities   []Abilities `json:"abilities"`
}

type Filter struct {
	Roles [] Role `json:"roles"`
	Name string `json:"name"`
}

type FilterRequest struct {
	Filter Filter `json:"filter"`
} 

type Response struct {
	Status int     `json:"status"`
	Data   []Agent `json:"data"`
}