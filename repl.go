package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Alb3G/pokedexcli/internal"
)

func startRepl(conf *internal.Config) {
	scanner := bufio.NewScanner(os.Stdin)

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
