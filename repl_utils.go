package main

import (
	"fmt"
	"os"
	"strings"
)

var supportedCommands map[string]CliCommand

func init() {
	supportedCommands = map[string]CliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCommand,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas from Pokemon World",
			callback:    mapCommand,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit(conf *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCommand(conf *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for key, value := range supportedCommands {
		fmt.Printf("%v: %v\n", key, value.description)
	}

	return nil
}

func mapCommand(conf *Config) error {

	return nil
}
