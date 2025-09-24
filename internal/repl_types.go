package internal

import (
	"sync"
	"time"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(conf *Config, args []string) error
}

type Config struct {
	PreviousUrl *string
	NextUrl     *string
	Client      Client
}

type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Pokemon struct {
	Id             int    `json:"id"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Name           string `json:"name"`

	Stats []struct {
		BaseStat int  `json:"base_stat"`
		Effort   int  `json:"effort"`
		Stat     Stat `json:"stat"`
	} `json:"stats"`

	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

type Stat struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Pokedex struct {
	Data map[string]Pokemon
}

type CatchProbability struct {
	Probability int
	IsCaught    bool
	CatchedMsg  string
}

type PokemonEncounter struct {
	Pokemon        Pokemon `json:"pokemon"`
	VersionDetails []any   `json:"version_details"`
}

type PokemonEncounterResp struct {
	Id                int                `json:"id"`
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type Cache struct {
	Elements map[string]CacheEntry
	Mutex    *sync.Mutex
}

type CacheEntry struct {
	CratedAt time.Time
	Val      []byte
}

const BASE_URL = "https://pokeapi.co/api/v2"
