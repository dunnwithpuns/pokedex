package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Pokedex > ")

	for scanner.Scan() {
		fmt.Println("Pokedex > ")
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
}

func initCommands() {
	type cliCommand struct {
		name        string
		description string
		callback    func() error
	}

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
	}
}
