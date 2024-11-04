package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasResponse, error) {
	endpoint := "location-area"
	fullURL := baseURL + endpoint

	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer response.Body.Close()

	if response.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasResponse := LocationAreasResponse{}
	if err = json.Unmarshal(data, &locationAreasResponse); err != nil {
		return LocationAreasResponse{}, err
	}

	return locationAreasResponse, nil
}
