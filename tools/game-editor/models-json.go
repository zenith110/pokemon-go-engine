package main

type PokemonJson struct {
	Species        string   `json:"species"`
	Level          int      `json:"level"`
	HP             int      `json:"hp"`
	Attack         int      `json:"attack"`
	Defense        int      `json:"defense"`
	Speed          int      `json:"speed"`
	SpecialAttack  int      `json:"specialAttack"`
	SpecialDefense int      `json:"specialDefense"`
	Moves          []string `json:"moves"`
	HeldItem       string   `json:"heldItem"`
}

type TrainerJson struct {
	Name      string        `json:"name"`
	Sprite    string        `json:"sprite"`
	Id        string        `json:"id"`
	Pokemons  []PokemonJson `json:"pokemons"`
	ClassType string        `toml:"classType"`
}
