package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetPokemon fetches details for a specific pokemon by name
func (c *Client) GetPokemon(name string) (PokemonResponse, error) {
	url := baseURL + "/pokemon/" + name

	// Check cache first
	if data, found := c.cache.Get(url); found {
		// fmt.Println("Cache hit! Using cached data for:", url)

		pokemonResp := PokemonResponse{}
		err := json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return PokemonResponse{}, err
		}
		return pokemonResp, nil
	}

	// fmt.Println("Cache miss! Fetching data from API:", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer resp.Body.Close()

	// Handle non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return PokemonResponse{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	// Add to cache
	c.cache.Add(url, dat)

	pokemonResp := PokemonResponse{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return PokemonResponse{}, err
	}

	return pokemonResp, nil
}
