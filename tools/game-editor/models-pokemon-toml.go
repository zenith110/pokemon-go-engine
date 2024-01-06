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
	Level int    `toml:"level"`
	Name  string `toml:"name"`
	ID    int    `toml:"id"`
}
type Types struct {
	Type1 string `toml:"type1"`
	Type2 string `toml:"type2"`
}
type AssetData struct {
	FrontSprite string `toml:"frontSprite"`
	BackSprite  string `toml:"backSprite"`
	Icon        string `toml:"icon"`
	Cry         string `toml:"cry"`
}
type LevelUpMoves struct {
	Level int    `toml:"level"`
	Name  string `toml:"name"`
	Type1 string `toml:"type1"`
	Type2 string `toml:"type2"`
}
type Pokemon struct {
	Name         string         `toml:"name"`
	ID           int            `toml:"id"`
	Stats        Stats          `toml:"stats"`
	Evolution    Evolution      `toml:"evolution"`
	BaseRate     int            `toml:"baseRate"`
	Types        Types          `toml:"types"`
	AssetData    AssetData      `toml:"assetData"`
	Ability1     []string       `toml:"ability1"`
	Ability2     []string       `toml:"ability2"`
	LevelUpMoves []LevelUpMoves `toml:"levelUpMoves"`
}

type PokemonTrainerEditor struct {
	Name           string
	FrontSprite    string
	BackSprite     string
	Icon           string
	HP             int
	Defense        int
	SpecialAttack  int
	Speed          int
	SpecialDefense int
	Attack         int
	Moves          []LevelUpMoves
}
