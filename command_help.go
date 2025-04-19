package main

import (
	"fmt"
	"github.com/datsun80zx/Go_Pokedex/internal/pokeapi"
)

func commandHelp(config *pokeapi.Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}