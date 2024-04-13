package models

type Maps struct {
	Map []Map `toml:"map"`
}
type MapPokemon struct {
	Species      string `toml:"species"`
	MaxLevel     string `toml:"maxLevel"`
	MinLevel     string `toml:"minLevel"`
	PriorityTile string `toml:"priorityTile"`
}
type Encounter struct {
	Pokemon MapPokemon `toml:"pokemon"`
}
type Map struct {
	Name        string    `toml:"name"`
	Filepath    string    `toml:"filepath"`
	Tilesetpath string    `toml:"tilesetpath"`
	Description string    `toml:"description"`
	TypeOfMap   string    `toml:"typeOfMap"`
	BgMusic     string    `toml:"bgMusic"`
	ID          string    `toml:"id"`
	Encounter   Encounter `toml:"encounter"`
}
