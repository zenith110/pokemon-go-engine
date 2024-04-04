package models

type MysteryGift struct {
	Name     string `toml:"Name"`
	GiftType string `toml:"GiftType"`
	FileName string `toml:"FileName"`
}
type PokemonGift struct {
	Name             string      `toml:"Name"`
	OriginalTrainer  string      `toml:"OriginalTrainer"`
	BeginningDate    string      `toml:"BeginningDate"`
	EndDate          string      `toml:"EndDate"`
	ID               string      `toml:"id"`
	Species          string      `toml:"species"`
	Types            []string    `toml:"types"`
	Abilities        []Abilities `toml:"abilities"`
	Moves            []Moves     `toml:"moves"`
	Stats            Stats       `toml:"stats"`
	AssetData        AssetData   `toml:"assetData"`
	Nickname         string      `toml:"Nickname"`
	LocationOfOrigin string      `toml:"LocationOfOrigin"`
	Shiny            bool        `toml:"Shiny"`
}
