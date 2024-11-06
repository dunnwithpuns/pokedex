package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		validCommands := getCommands()

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)

		inputParameter := ""

		if len(cleaned) == 0 {
			continue
		}

		if len(cleaned) > 1 {
			inputParameter = cleaned[1]
		}

		inputCommand := cleaned[0]

		command, ok := validCommands[inputCommand]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		} else if command.name != "explore" && inputParameter != "" {
			fmt.Printf("The %v command does not accept any parameters", command.name)
			continue
		}

		if err := command.callback(cfg, inputParameter); err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)

	if len(words) > 1 {

	}
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display 20 new location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 locations provided by map",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "List all pokemon in the described area",
			callback:    commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}

}
