package main

type TrainerToml struct {
	Trainers []Trainer `toml:"trainer"`
}
type Pokemons struct {
	Species  string   `toml:"species"`
	Level    int      `toml:"level"`
	Moves    []string `toml:"moves"`
	HeldItem string   `toml:"heldItem"`
}
type Trainer struct {
	Name      string     `toml:"name"`
	Sprite    string     `toml:"sprite"`
	ID        string     `toml:"id"`
	Pokemons  []Pokemons `toml:"pokemons"`
	ClassType string     `toml:"classType"`
}
