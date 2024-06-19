package models

type EncounterEnv struct {
	Name       string  `toml:"name"`
	Percentage float64 `toml:"percentage"`
}

type MapPokemon struct {
	Species      string       `toml:"species"`
	MaxLevel     string       `toml:"maxLevel"`
	MinLevel     string       `toml:"minLevel"`
	PriorityTile EncounterEnv `toml:"encounterEnvironment"`
}

type Encounter struct {
	Pokemon MapPokemon `toml:"pokemon"`
}
type Event struct {
	XPos        string `toml:"xPosition"`
	YPos        string `toml:"yPosition"`
	ScriptName  string `toml:"scriptName"`
	OverworldId string `toml:"overworldId"`
}

type MapToml struct {
	Name        string      `toml:"name"`
	Tilesetpath string      `toml:"tilesetpath"`
	TmxFilePath string      `toml:"tmxFilePath"`
	Description string      `toml:"description"`
	TypeOfMap   string      `toml:"typeOfMap"`
	BgMusic     string      `toml:"bgMusic"`
	ID          string      `toml:"id"`
	Encounters  []Encounter `toml:"encounter"`
	Events      []Event     `toml:"event"`
}
