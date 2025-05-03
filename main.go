package main

import (
	"time"

	"github.com/datsun80zx/Go_Pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.PokemonResponse),
	}

	startRepl(cfg)
}
