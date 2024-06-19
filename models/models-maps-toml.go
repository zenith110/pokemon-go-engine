package models

type MapData struct {
	Maps []Map `toml:"map"`
}

type Map struct {
	Name     string `toml:"name"`
	Filepath string `toml:"filepath"`
	ID       string `toml:"id"`
}
