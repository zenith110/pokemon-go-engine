package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/gofrs/uuid"
	"github.com/pelletier/go-toml/v2"
)

type Pokemon struct {
	species  string   `toml:"species"`
	level    int      `toml:"level"`
	moves    []string `toml:"moves"`
	heldItem string   `toml:"heldItem"`
}

type Trainer struct {
	name      string `toml:"name"`
	sprite    string `toml:"sprite"`
	id        string `toml:"id"`
	pokemon1  Pokemon
	pokemon2  Pokemon
	pokemon3  Pokemon
	pokemon4  Pokemon
	pokemon5  Pokemon
	pokemon6  Pokemon
	classType string `toml:"classType"`
}

type Trainers struct {
	Trainer struct {
		name      string `toml:"name"`
		sprite    string `toml:"sprite"`
		id        string `toml:"id"`
		pokemon1  Pokemon
		pokemon2  Pokemon
		pokemon3  Pokemon
		pokemon4  Pokemon
		pokemon5  Pokemon
		pokemon6  Pokemon
		classType string `toml:"classType"`
	}
}

type Config struct {
	dataDirectory    string `toml:"dataDirectory"`
	updateFromGithub bool   `toml:"updateFromGithub"`
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func SetupConfig() Config {
	file, err := os.Open("settings.toml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var config Config
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal(bytes, &config)
	if err != nil {
		panic(err)
	}
	return config
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	file, err := os.Open("settings.toml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

}

func (a *App) LoadTrainerToml() Trainers {
	config := SetupConfig()
	file, err := os.Open(fmt.Sprintf("%s/trainers.toml", config.dataDirectory))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var trainers Trainers
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal(bytes, &trainers)
	if err != nil {
		panic(err)
	}
	return trainers
}

func (a *App) CreateTrainerData() {
	config := SetupConfig()
	var pokemon1 = Pokemon{
		species:  "Blastoise",
		level:    100,
		heldItem: "Oran Berry",
	}
	var trainer = Trainer{
		name:      "",
		sprite:    "",
		id:        uuid.NewV7(),
		pokemon1:  pokemon1,
		classType: "",
	}
	data, err := toml.Marshal(trainer)
	if err != nil {
		panic(err)
	}

	// Write the encoded data to a file
	err = os.WriteFile(fmt.Sprintf("%s/trainers.toml", config.dataDirectory), data, 0644)
	if err != nil {
		panic(err)
	}
}
