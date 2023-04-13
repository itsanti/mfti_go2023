package main

type Pokemon struct {
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	ID             int    `json:"id"`
	IsDefault      bool   `json:"is_default"`
	Name           string `json:"name"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
}

type APIclient interface {
	GetPokemons() []Pokemon
}

func GetAPIclient(mock bool) APIclient {
	var a APIclient
	return a
}
