package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetLocationArea fetches details for a specific location area by name
func (c *Client) GetLocationArea(name string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area/" + name

	// Check cache first
	if data, found := c.cache.Get(url); found {
		// fmt.Println("Cache hit! Using cached data for:", url)

		locationResp := LocationAreaResponse{}
		err := json.Unmarshal(data, &locationResp)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return locationResp, nil
	}

	// fmt.Println("Cache miss! Fetching data from API:", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	// Handle non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return LocationAreaResponse{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	// Add to cache
	c.cache.Add(url, dat)

	locationResp := LocationAreaResponse{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationResp, nil
}
