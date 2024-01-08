package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/pelletier/go-toml/v2"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func SetupConfig() OptionsConfig {
	file, err := os.Open("settings.toml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var config OptionsConfig
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

// func (a *App) LoadTrainerToml() Trainer {
// 	file, err := os.Open("trainers.toml")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	var trainers Trainer
// 	bytes, err := io.ReadAll(file)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = toml.Unmarshal(bytes, &trainers)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return trainers
// }

func (a *App) CreateTrainerData(trainerJson TrainerJson) {

	var pokemons []Pokemons
	for index := range trainerJson.Pokemons {
		pokemon := Pokemons{
			Species:        trainerJson.Pokemons[index].Species,
			Level:          trainerJson.Pokemons[index].Level,
			Moves:          trainerJson.Pokemons[index].Moves,
			HeldItem:       trainerJson.Pokemons[index].HeldItem,
			HP:             trainerJson.Pokemons[index].HP,
			Defense:        trainerJson.Pokemons[index].Defense,
			Attack:         trainerJson.Pokemons[index].Attack,
			SpecialAttack:  trainerJson.Pokemons[index].SpecialAttack,
			SpecialDefense: trainerJson.Pokemons[index].SpecialDefense,
			Speed:          trainerJson.Pokemons[index].Speed,
		}
		pokemons = append(pokemons, pokemon)
	}
	var trainers []Trainers

	trainer := Trainers{
		Name:      trainerJson.Name,
		Sprite:    trainerJson.Sprite,
		ID:        trainerJson.Id,
		Pokemons:  pokemons,
		ClassType: trainerJson.ClassType,
	}
	fmt.Print(trainer)
	trainers = append(trainers, trainer)
	trainerConfig := TrainerToml{
		Trainers: trainers,
	}
	data, err := toml.Marshal(trainerConfig)
	if err != nil {
		panic(fmt.Errorf("Error had occured while creating trainer data!\n%v", err))
	}

	// Write the encoded data to a file
	f, err := os.OpenFile("trainers.toml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := f.Write(data); err != nil {
		panic(err)
	}
}

func (a *App) ParsePokemonData() []PokemonTrainerEditor {
	file, err := os.Open("pokemon.toml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var pokemons PokemonToml

	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = toml.Unmarshal(b, &pokemons)
	if err != nil {
		panic(err)
	}
	var trainerEditorPokemons []PokemonTrainerEditor
	for pokemon := range pokemons.Pokemon {
		trainerEditorPokemon := PokemonTrainerEditor{
			Name:           pokemons.Pokemon[pokemon].Name,
			FrontSprite:    pokemons.Pokemon[pokemon].AssetData.FrontSprite,
			BackSprite:     pokemons.Pokemon[pokemon].AssetData.BackSprite,
			Icon:           pokemons.Pokemon[pokemon].AssetData.Icon,
			HP:             pokemons.Pokemon[pokemon].Stats.Hp,
			Defense:        pokemons.Pokemon[pokemon].Stats.Defense,
			SpecialAttack:  pokemons.Pokemon[pokemon].Stats.SpecialAtk,
			Speed:          pokemons.Pokemon[pokemon].Stats.Speed,
			SpecialDefense: pokemons.Pokemon[pokemon].Stats.SpecialDef,
			Moves:          pokemons.Pokemon[pokemon].LevelUpMoves,
			Attack:         pokemons.Pokemon[pokemon].Stats.Attack,
		}
		trainerEditorPokemons = append(trainerEditorPokemons, trainerEditorPokemon)
	}
	return trainerEditorPokemons
}

func (a *App) ParseTrainerClass() TrainerClasses {
	file, err := os.Open("trainerclasses.toml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var trainerclasses TrainerClasses
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal(bytes, &trainerclasses)
	if err != nil {
		panic(err)
	}

	return trainerclasses
}

func (a *App) ParseHeldItems() []HeldItem {
	file, err := os.Open("helditems.toml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var heldItems HeldItemToml

	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = toml.Unmarshal(b, &heldItems)
	if err != nil {
		panic(err)
	}
	var heldItemsData []HeldItem
	for heldItem := range heldItems.HeldItems {
		heldItemData := HeldItem{
			Name: heldItems.HeldItems[heldItem].Name,
		}
		heldItemsData = append(heldItemsData, heldItemData)
	}
	return heldItemsData
}
