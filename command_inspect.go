package main

import (
	"fmt"
)

func commandInspect(cfg *config, pokemonName string) error {
	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		fmt.Println("You have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")

	for i := range pokemon.Stats {
		fmt.Printf("	-%v: %v\n", pokemon.Stats[i].Stat.Name, pokemon.Stats[i].BaseStat)
	}

	fmt.Println("Types:")

	for i := range pokemon.Types {
		fmt.Printf("	- %s\n", pokemon.Types[i].Type.Name)
	}

	return nil
}
