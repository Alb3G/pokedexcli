package internal

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageUrl *string) (LocationArea, error) {
	url := BASE_URL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	// Check if cache contains any data before making the request
	if data, ok := c.cache.Get(url); ok {
		var locationArea LocationArea
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	// Make request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	// Add response to Cache
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	var locationArea LocationArea
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, nil
	}

	c.cache.Add(url, data)
	return locationArea, nil
}

func (c *Client) GetPokemonsFromLocation(locationName *string) (PokemonEncounterResp, error) {
	url := BASE_URL + "/location-area/" + *locationName

	if data, ok := c.cache.Get(url); ok {
		var pokemonEncounters PokemonEncounterResp
		err := json.Unmarshal(data, &pokemonEncounters)
		if err != nil {
			return PokemonEncounterResp{}, err
		}
		return pokemonEncounters, nil
	}

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonEncounterResp{}, err
	}

	// Get Response
	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonEncounterResp{}, err
	}
	defer res.Body.Close()

	// Process data
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonEncounterResp{}, err
	}

	var pokemonEncounters PokemonEncounterResp
	err = json.Unmarshal(data, &pokemonEncounters)
	if err != nil {
		return PokemonEncounterResp{}, err
	}

	// Add data to cache
	c.cache.Add(url, data)

	return pokemonEncounters, nil
}

func (c *Client) GetPokemonByName(name *string) (Pokemon, error) {
	url := BASE_URL + "/pokemon/" + *name

	// Check Cache first
	if data, ok := c.cache.Get(url); ok {
		var pokemon Pokemon
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	return pokemon, nil
}
