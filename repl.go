package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var registry map[string]cliCommand

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n\n")
	for key, value := range registry {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func init() {
	registry = map[string]cliCommand{
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
	}
}

func startRepl() {

	// so i need to set up a new scanner
	userInput := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		userInput.Scan()
		words := cleanInput(userInput.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := registry[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(strings.TrimSpace(text))
	words := strings.Fields(lowered)
	return words
}
