package main

type PokemonJson struct {
	Species  string   `json:"species"`
	Level    int      `json:"level"`
	Moves    []string `json:"moves"`
	HeldItem string   `json:"heldItem"`
}

type TrainerJson struct {
	Name      string        `json:"name"`
	Sprite    string        `json:"sprite"`
	Id        string        `json:"id"`
	Pokemons  []PokemonJson `json:"pokemons"`
	ClassType string        `toml:"classType"`
}
