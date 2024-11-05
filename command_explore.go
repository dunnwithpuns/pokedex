package main

import (
	"fmt"
)

func commandExplore(cfg *config, commandParameter string) error {
	response, err := cfg.pokeapiClient.ExploreArea(commandParameter)
	if err != nil {
		return err
	}

	fmt.Println("Pokemon in the area:")

	for _, pokemon := range response.PokemonEncounters {
		fmt.Printf("- %v", pokemon)
	}

	return nil
}
