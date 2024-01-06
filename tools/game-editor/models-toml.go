package main

type OptionsConfig struct {
	dataDirectory    string `toml:"dataDirectory"`
	updateFromGithub bool   `toml:"updateFromGithub"`
}
