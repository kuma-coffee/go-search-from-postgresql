package entity

type Post struct {
	Ndex       int      `json:"ndex" gorm:"primaryKey"`
	Pokemon    string   `json:"pokemon" gorm:"type:varchar(255);not null"`
	PokemonURL string   `json:"pokemon_url" gorm:"type:varchar(255);not null"`
	Type       []string `json:"type" gorm:"type:varchar(255);"`
}
