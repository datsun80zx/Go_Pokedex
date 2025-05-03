package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("please provide a Pokemon name")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	// Get Pokemon data
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("failed to find pokemon: %s", err)
	}

	// Calculate catch chance based on base experience
	// Higher base experience = harder to catch
	catchChance := calculateCatchChance(pokemon.BaseExperience)

	// Random number between 0 and 1
	// As of Go 1.20, the global rand is automatically seeded, so this is properly random
	randNum := rand.Float64()

	// If random number is less than catch chance, the Pokemon is caught
	caught := randNum <= catchChance

	if caught {
		fmt.Printf("%s was caught!\n", pokemonName)
		// Add to Pokedex
		cfg.caughtPokemon[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

// calculateCatchChance returns a probability between 0.1 and 0.9
// based on the Pokemon's base experience
func calculateCatchChance(baseExp int) float64 {
	// Base experience typically ranges from 40 to 600
	// Map this to a probability between 0.1 and 0.9
	// Lower base experience = higher catch chance

	// If base experience is very low or not set, use a default
	if baseExp <= 0 {
		baseExp = 100
	}

	// Max base experience we'll consider (higher will still use this cap)
	maxBaseExp := 600.0

	// Limit base experience to our max
	cappedBaseExp := float64(baseExp)
	if cappedBaseExp > maxBaseExp {
		cappedBaseExp = maxBaseExp
	}

	// Calculate catch chance: 0.9 for lowest base exp, 0.1 for highest
	// Linear interpolation between these values
	catchChance := 0.9 - 0.8*(cappedBaseExp/maxBaseExp)

	return catchChance
}
