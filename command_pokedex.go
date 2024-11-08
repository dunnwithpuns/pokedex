package main

import (
	"fmt"
)

func commandPokedex(cfg *config, inputParam string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("--- Pokedex Empty ---")
		return nil
	}

	fmt.Println("Your Pokedex:")

	for _, pokemon := range cfg.pokedex {
		fmt.Println("- " + pokemon.Name)
	}
	return nil

}
