package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Config struct {
	NextURL      *string
	PreviousURL  *string
}

type LocationAreaResponse struct {
	Count int `json:"count"` 
	Next *string `json:"next"` 
	Previous *string `json:"previous"` 
	Results []struct { 
		Name string `json:"name"` 
		URL string `json:"url"` 
	} `json:"results"`
}

func GetLocationAreas(url string) (*LocationAreaResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var locations LocationAreaResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return nil, err
	}
	return &locations, nil
}