package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Wait for user input
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

		if err := cliCommandStr.callback(); err != nil {
			fmt.Println(err)
		}
	}
}
