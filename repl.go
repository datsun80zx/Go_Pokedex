package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

		fmt.Printf("Your command was: %s\n", commandName)
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(strings.TrimSpace(text))
	words := strings.Fields(lowered)
	return words
}
