package main

import (
	"fmt"
)

func commandMap(cfg *config, commandParameter string) error {
	response, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range response.Results {
		fmt.Println("- ", area.Name)
	}

	cfg.nextLocationAreaURL = response.Next
	cfg.previousLocationAreaURL = response.Previous

	return nil
}

func commandMapB(cfg *config, commandParameter string) error {
	if cfg.previousLocationAreaURL == nil {
		return fmt.Errorf("No previous page")
	}

	response, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}

	for _, area := range response.Results {
		fmt.Println("- ", area.Name)
	}

	cfg.nextLocationAreaURL = response.Next
	cfg.previousLocationAreaURL = response.Previous

	return nil
}
