package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreArea(area string) (AreaResponse, error) {
	endpoint := "location-area/"
	fullURL := baseURL + endpoint + area

	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit")
		areaResponse := AreaResponse{}
		if err := json.Unmarshal(data, &areaResponse); err != nil {
			return AreaResponse{}, err
		}
		return areaResponse, nil
	}
	fmt.Println("cache miss")

	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return AreaResponse{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return AreaResponse{}, err
	}

	defer response.Body.Close()

	if response.StatusCode > 399 {
		return AreaResponse{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return AreaResponse{}, err
	}

	areaResponse := AreaResponse{}
	if err = json.Unmarshal(data, &areaResponse); err != nil {
		return AreaResponse{}, err
	}

	c.cache.Add(fullURL, data)

	return areaResponse, nil
}
