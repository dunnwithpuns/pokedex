package main

import (
	"fmt"
	"github.com/dunnwithpuns/pokedex/internal/pokeapi"
	"log"
)

func commandMap() error {
	pokeapiClient := pokeapi.NewClient()

	response, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	return nil
}
