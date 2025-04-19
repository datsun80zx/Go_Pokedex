package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/datsun80zx/Go_Pokedex/internal/pokeapi"
)


func startRepl() {

	config := &pokeapi.Config{
		NextURL: nil,
		PreviousURL: nil,
	}

	userInput := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		userInput.Scan()

		words := cleanInput(userInput.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(strings.TrimSpace(text))
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *pokeapi.Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map", 
			description: "Displays the names of the next 20 location areas in the Pokemon world",
			callback:	 commandMap,
		},
		"mapb": {
			name:        "mapb", 
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:	 commandMapb,
		},
	}
}


