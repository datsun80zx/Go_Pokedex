package main

import (
	"fmt"
	"github.com/datsun80zx/Go_Pokedex/internal/pokeapi"
)

func commandMap(config *pokeapi.Config) error {
	apiURL := "https://pokeapi.co/api/v2/location-area"
	if config.NextURL != nil && *config.NextURL != "" {
		apiURL = *config.NextURL
	}

	locationData, err := pokeapi.GetLocationAreas(apiURL)
	if err != nil {
		fmt.Errorf("Error getting location data: %w", err)
	}
	if locationData.Next != nil {
		config.NextURL = locationData.Next
	} else {
		config.NextURL = nil
	}

	if locationData.Previous != nil {
		config.PreviousURL = locationData.Previous
	} else {
		config.PreviousURL = nil
	}

	if locationData.Results == nil {
		return fmt.Errorf("no location data available")
	}
	for _, loc := range locationData.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(config *pokeapi.Config) error {
	if config.PreviousURL == nil || *config.PreviousURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	apiURL := *config.PreviousURL

	locationData, err := pokeapi.GetLocationAreas(apiURL)
	if err != nil {
		fmt.Errorf("Error getting location data: %w", err)
	}
	if locationData.Next != nil {
		config.NextURL = locationData.Next
	} else {
		config.NextURL = nil
	}

	if locationData.Previous != nil {
		config.PreviousURL = locationData.Previous
	} else {
		config.PreviousURL = nil
	}

	if locationData.Results == nil {
		return fmt.Errorf("no location data available")
	}
	for _, loc := range locationData.Results {
		fmt.Println(loc.Name)
	}
	return nil
}