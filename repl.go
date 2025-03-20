package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowered := strings.ToLower(strings.TrimSpace(text))
	words := strings.Fields(lowered)
	return words
}
