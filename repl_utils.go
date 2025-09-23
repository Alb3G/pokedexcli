package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Alb3G/pokedexcli/internal"
)

var supportedCommands map[string]internal.CliCommand

func init() {
	supportedCommands = map[string]internal.CliCommand{
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
		"mapb": {
			Name:        "mapb",
			Description: "Displays the 20 previous location areas from Pokemon World",
			Callback:    mapBackCommand,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays all the pokemons found in an especific location",
			Callback:    exploreCommand,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit(conf *internal.Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCommand(conf *internal.Config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for key, value := range supportedCommands {
		fmt.Printf("%v: %v\n", key, value.Description)
	}

	return nil
}

func mapCommand(conf *internal.Config, args []string) error {
	locationArea, err := conf.Client.GetLocationAreas(conf.NextUrl)
	if err != nil {
		return err
	}

	conf.NextUrl = locationArea.Next
	conf.PreviousUrl = locationArea.Previous

	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func mapBackCommand(conf *internal.Config, args []string) error {
	if conf.PreviousUrl == nil {
		return errors.New("you're on the first page")
	}

	locationArea, err := conf.Client.GetLocationAreas(conf.PreviousUrl)
	if err != nil {
		return err
	}

	conf.NextUrl = locationArea.Next
	conf.PreviousUrl = locationArea.Previous
	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func exploreCommand(conf *internal.Config, args []string) error {
	fmt.Printf("Exploring %s...\n", args[0])

	res, err := conf.Client.GetPokemonsFromLocation(&args[0])
	if err != nil {
		return errors.New("unknown location name")
	}

	if len(res.PokemonEncounters) == 0 {
		fmt.Println("No pokemons encountered in this area")
		return nil
	}

	for _, encounter := range res.PokemonEncounters {
		fmt.Print(" - ")
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
