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

		input := cleanInput(scanner.Text())

		command, isSupported := supportedCommands[input[0]]

		if !isSupported {
			fmt.Println("Unknown command")
			continue
		}

		err := command.Callback(conf, input[1:])

		if err != nil {
			fmt.Println(err)
		}
	}
}
