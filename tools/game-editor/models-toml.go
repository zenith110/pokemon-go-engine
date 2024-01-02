package main

type OptionsConfig struct {
	dataDirectory    string `toml:"dataDirectory"`
	updateFromGithub bool   `toml:"updateFromGithub"`
}

type TrainerClass struct {
	name  string `toml:"name"`
	music string `toml:"music"`
}
