package main

type PokemonToml struct {
	Pokemon []Pokemon `toml:"pokemon"`
}
type Stats struct {
	Hp         int `toml:"hp"`
	Speed      int `toml:"speed"`
	Defense    int `toml:"defense"`
	SpecialAtk int `toml:"specialAtk"`
	SpecialDef int `toml:"specialDef"`
	Attack     int `toml:"attack"`
}
type Evolution struct {
	Name   string   `toml:"name"`
	ID     int      `toml:"id"`
	Method []string `toml:"method"`
}

type AssetData struct {
	FrontSprite string `toml:"frontSprite"`
	BackSprite  string `toml:"backSprite"`
	Icon        string `toml:"icon"`
	Cry         string `toml:"cry"`
}
type Ability struct {
	Name     string `toml:"name"`
	IsHidden bool   `toml:"isHidden"`
}
type Move struct {
	Level  int    `toml:"level"`
	Name   string `toml:"name"`
	Method string `toml:"method"`
}
type Pokemon struct {
	Name       string      `toml:"name"`
	ID         int         `toml:"id"`
	Stats      Stats       `toml:"stats"`
	Evolutions []Evolution `toml:"evolutions"`
	BaseRate   int         `toml:"baseRate"`
	Types      []string    `toml:"types"`
	AssetData  AssetData   `toml:"assetData"`
	Abilities  []Ability   `toml:"ability"`
	Moves      []Move      `toml:"moves"`
	DexEntry   string      `toml:"dexEntry"`
}
