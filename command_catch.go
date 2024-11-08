package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, name string) error {

	response, err := cfg.pokeapiClient.CatchPokemon(name)
	if err != nil {
		return err
	}

	chanceToCatch := rand.Intn(response.BaseExperience)
	success := false

	if chanceToCatch > int(response.BaseExperience/2) {
		success = true
	}

	fmt.Printf("- Throwing a Pokeball at %s. . .\n", name)
	time.Sleep(time.Second)

	if !success {
		fmt.Printf("** %s escaped! **\n", name)
		return nil
	}

	fmt.Printf("-- %s was caught! --\n", name)
	cfg.pokedex[response.Name] = response
	return nil
}
