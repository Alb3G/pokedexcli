package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Alb3G/pokedexcli/internal/types"
)

func main() {
	// Wait for user input
	scanner := bufio.NewScanner(os.Stdin)

	conf := &types.Config{
		PreviousUrl: "",
		NextUrl:     "",
	}

	for {
		fmt.Print("pokedex > ")

		scanner.Scan()

		command := cleanInput(scanner.Text())

		cliCommandStr, isSupported := supportedCommands[command[0]]

		if !isSupported {
			fmt.Println("Unknown command")
			continue
		}

		err := cliCommandStr.Callback(conf)

		if err != nil {
			fmt.Println(err)
		}
	}
}
