package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(name string) (PokemonResponse, error) {
	endpoint := "pokemon/"
	fullURL := baseURL + endpoint + name

	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit")
		pokemon := PokemonResponse{}
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return PokemonResponse{}, err
		}
		return pokemon, nil
	}
	fmt.Println("cache miss")

	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return PokemonResponse{}, err
	}

	defer response.Body.Close()

	if response.StatusCode > 399 {
		return PokemonResponse{}, fmt.Errorf("bad status code: ", response.StatusCode)
	}

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	pokemon := PokemonResponse{}
	if err = json.Unmarshal(data, &pokemon); err != nil {
		return PokemonResponse{}, err
	}

	return pokemon, nil
}
