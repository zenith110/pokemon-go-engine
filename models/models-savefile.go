package models

type LastSavedLocation struct {
	XPos     string `toml:"xpos"`
	YPos     string `toml:"ypos"`
	CityName string `toml:"cityname"`
}
type SaveFilePokemon struct {
	HP             string `toml:"hp"`
	Defense        string `toml:"defense"`
	Speed          string `toml:"speed"`
	SpecialAttack  string `toml:"specialAttack"`
	SpecialDefense string `toml:"specialDefense"`
	Moves          []Move `toml:"moves"`
	Level          string `toml:"level"`
	Species        string `toml:"species"`
	Ability        string `toml:"ability"`
}

type SaveFile struct {
	Name               string            `toml:"name"`
	LastSavedDate      string            `toml:"lastSavedDate"`
	Money              string            `toml:"money"`
	CurrentBadgeAmount string            `toml:"currentbadgeamount"`
	SavedLocation      LastSavedLocation `toml:"savedLocation"`
	CurrentPlayTime    string            `toml:"currentPlayTime"`
	Pokemons           []SaveFilePokemon `toml:"pokemons"`
	LastQuest          string            `toml:"lastQuest"`
}
