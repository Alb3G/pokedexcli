package main

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/Alb3G/pokedexcli/internal"
)

func getDifficulty(baseExperience int) float64 {
	switch {
	case baseExperience <= 50:
		return 0.4
	case baseExperience <= 110:
		return 0.6
	default:
		return 0.8
	}
}

func getCatchMessage(name string, isCaught bool) string {
	if isCaught {
		return fmt.Sprintf("%s was caught!", name)
	}
	return fmt.Sprintf("%s escaped!", name)
}

func catchPokemonProbability(name string, baseExperience int) internal.CatchProbability {
	difficulty := getDifficulty(baseExperience)
	userProb := rand.IntN(baseExperience)
	catchSuccessRate := int(float64(baseExperience) * difficulty)
	isCaught := userProb < catchSuccessRate
	catchMsg := getCatchMessage(name, isCaught)

	return internal.CatchProbability{
		Probability: userProb,
		IsCaught:    isCaught,
		CatchedMsg:  catchMsg,
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func showPokemonFromPokedex(pokemon internal.Pokemon) {
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")

	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, types := range pokemon.Types {
		fmt.Printf(" - %v\n", types.Type.Name)
	}
}

func printPokedex(pokedex internal.Pokedex) {
	for _, v := range pokedex.Data {
		fmt.Print(" - ")
		fmt.Println(v.Name)
	}
}
