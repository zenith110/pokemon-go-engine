package main

import (
	"log"

	"github.com/joho/godotenv"
	Core "github.com/zenith110/pokemon-go-engine/core"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Core.CreateGame("Pokemon Go Engine", 640, 400)
}
