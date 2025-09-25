package main

import (
	"errors"
	"fmt"
	"os"

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
		"catch": {
			Name:        "catch",
			Description: "Catch a pokemon specified by argument",
			Callback:    catchCommand,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Display a pokemon stats",
			Callback:    inspectCommand,
		},
		"pokedex": {
			Name:        "catch",
			Description: "Catch a pokemon specified by argument",
			Callback:    pokedexCommand,
		},
	}
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

/*
For future features maybe we can check wether the pokemon is in the area that the
the user is exploring or not in order to allow the catch action
*/
func catchCommand(conf *internal.Config, args []string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	pokemon, err := conf.Client.GetPokemonByName(&args[0])
	if err != nil {
		return errors.New("unknown pokemon name")
	}

	cp := catchPokemonProbability(pokemon.Name, pokemon.BaseExperience)

	if !cp.IsCaught {
		return errors.New(cp.CatchedMsg)
	}

	// Add pokemon to pokedex
	internal.GlobalPokedex.AddPokemon(pokemon)

	fmt.Println(cp.CatchedMsg)

	return nil
}

func inspectCommand(conf *internal.Config, args []string) error {
	pokemon, exists := internal.GlobalPokedex.Data[args[0]]
	if !exists {
		return errors.New("you have not caught that pokemon")
	}

	showPokemonFromPokedex(pokemon)

	return nil
}

func pokedexCommand(conf *internal.Config, args []string) error {
	fmt.Println("Your Pokedex:")
	printPokedex(*internal.GlobalPokedex)
	return nil
}
