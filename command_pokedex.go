package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("Your Pokedex is empty. Try catching some Pokemon!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s", pokemon.Name)
		// fmt.Printf("   Base Experience: %d\n", pokemon.BaseExperience)
		// fmt.Printf("   Height: %d\n", pokemon.Height)
		// fmt.Printf("   Weight: %d\n", pokemon.Weight)

		// // Print types
		// fmt.Print("   Types: ")
		// for i, t := range pokemon.Types {
		// 	if i > 0 {
		// 		fmt.Print(", ")
		// 	}
		// 	fmt.Print(t.Type.Name)
		// }
		// fmt.Println()

		// // Print stats
		// fmt.Println("   Stats:")
		// for _, stat := range pokemon.Stats {
		// 	fmt.Printf("     %s: %d\n", stat.Stat.Name, stat.BaseStat)
		// }
		fmt.Println()
	}

	return nil
}
