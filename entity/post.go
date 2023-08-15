package entity

import "github.com/lib/pq"

type Post struct {
	Ndex       string         `json:"ndex" gorm:"primaryKey"`
	Pokemon    string         `json:"pokemon"`
	PokemonURL string         `json:"pokemon_url"`
	Type       pq.StringArray `json:"type" gorm:"type:text[]"`
}
