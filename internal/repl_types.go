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
	BaseExperience int    `json:"base_experience"`
	Id             int    `json:"id"`
	Name           string `json:"name"`
}

type CatchProbability struct {
	Probability int
	IsCatched   bool
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
