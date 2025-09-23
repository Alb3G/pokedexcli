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
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit(conf *internal.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCommand(conf *internal.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for key, value := range supportedCommands {
		fmt.Printf("%v: %v\n", key, value.Description)
	}

	return nil
}

func mapCommand(conf *internal.Config) error {
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

func mapBackCommand(conf *internal.Config) error {
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
