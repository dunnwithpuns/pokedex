package main

import "fmt"

func commandHelp() error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println("")

	validCommands := getCommands()

	for _, command := range validCommands {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}

	fmt.Println("")

	return nil
}
