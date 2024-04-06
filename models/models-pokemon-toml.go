package models

type PokemonToml struct {
	Pokemon []Pokemon `toml:"pokemon"`
}
type Abilities struct {
	Name     string `toml:"name"`
	IsHidden bool   `toml:"isHidden"`
}
type Moves struct {
	Name   string `toml:"name"`
	Level  int    `toml:"level"`
	Method string `toml:"method"`
}
type Evolutions struct {
	Name    string   `toml:"name"`
	Methods []string `toml:"methods"`
	ID      string   `toml:"id"`
}
type Stats struct {
	Hp             int `toml:"hp"`
	Attack         int `toml:"attack"`
	Defense        int `toml:"defense"`
	SpecialAttack  int `toml:"special-attack"`
	SpecialDefense int `toml:"special-defense"`
	Speed          int `toml:"speed"`
}
type AssetData struct {
	Front      string `toml:"front"`
	Back       string `toml:"back"`
	ShinyFront string `toml:"shiny_front"`
	ShinyBack  string `toml:"shiny_back"`
	Icon       string `toml:"icon"`
}
type Pokemon struct {
	ID         string       `toml:"id"`
	Species    string       `toml:"species"`
	Types      []string     `toml:"types"`
	DexEntry   string       `toml:"dexEntry"`
	Abilities  []Abilities  `toml:"abilities"`
	Moves      []Moves      `toml:"moves"`
	Evolutions []Evolutions `toml:"evolutions"`
	Stats      Stats        `toml:"stats"`
	AssetData  AssetData    `toml:"assetData"`
}

type AllMoves struct {
	Move []Move `toml:"move"`
}
type Descriptions struct {
	Description string `toml:"description"`
}
type Effects struct {
	EffectText string `toml:"effect_text"`
}
type Move struct {
	Name         string         `toml:"name"`
	Power        int            `toml:"Power"`
	Pp           int            `toml:"pp"`
	Accuracy     int            `toml:"accuracy"`
	Type         string         `toml:"type"`
	KindOfMove   string         `toml:"kind_of_move"`
	ID           int            `toml:"id"`
	Descriptions []Descriptions `toml:"descriptions"`
	Effects      []Effects      `toml:"effects"`
}
