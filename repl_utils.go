package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Alb3G/pokedexcli/api"
	"github.com/Alb3G/pokedexcli/types"
)

var supportedCommands map[string]types.CliCommand

func init() {
	supportedCommands = map[string]types.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    helpCommand,
		},
		"map": {
			Name:        "map",
			Description: "Displays 20 location areas from Pokemon World",
			Callback:    mapCommand,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit(conf *types.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCommand(conf *types.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for key, value := range supportedCommands {
		fmt.Printf("%v: %v\n", key, value.Description)
	}

	return nil
}

func mapCommand(conf *types.Config) error {
	var baseUrl string
	if conf.NextUrl == "" {
		baseUrl = "https://pokeapi.co/api/v2/location-area/"
	} else {
		baseUrl = conf.NextUrl
	}

	locationArea, err := api.GetLocationAreas(baseUrl)
	if err != nil {
		return err
	}

	conf.NextUrl = locationArea.Next
	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}

	return nil
}
