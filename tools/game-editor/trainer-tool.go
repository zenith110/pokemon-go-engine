package main

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

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
	f, err := os.OpenFile(fmt.Sprintf("%s/toml/trainers.toml", a.dataDirectory.DataDirectory), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := f.Write(data); err != nil {
		panic(err)
	}
}
