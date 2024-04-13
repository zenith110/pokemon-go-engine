package models

type MysteryGift struct {
	Name     string `toml:"Name"`
	GiftType string `toml:"GiftType"`
	FileName string `toml:"FileName"`
}
type PokemonGift struct {
	Name             string      `toml:"name"`
	OriginalTrainer  string      `toml:"originalTrainer"`
	BeginningDate    string      `toml:"beginningDate"`
	EndDate          string      `toml:"endDate"`
	ID               string      `toml:"id"`
	Species          string      `toml:"species"`
	Types            []string    `toml:"types"`
	Abilities        []Abilities `toml:"abilities"`
	Moves            []Moves     `toml:"moves"`
	Stats            Stats       `toml:"stats"`
	AssetData        AssetData   `toml:"assetData"`
	Nickname         string      `toml:"nickname"`
	LocationOfOrigin string      `toml:"locationOfOrigin"`
	Shiny            bool        `toml:"shiny"`
}
