package pokeapi

import (
	"fmt"
	"net/http"
	"encoding/json"
)

const baseUrl = "https://pokeapi.co/api/v2"

type ResLocationArea struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationArea(pageUrl *string) (ResLocationArea, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	res, err := http.Get(url)
	if err != nil {
		return ResLocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Errorf(
			"Response failed with status code: %d and\nbody: %s\n",
			res.StatusCode,
		)
	}

	var locationArea ResLocationArea
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationArea); err != nil {
		return ResLocationArea{}, err
	}

	return locationArea, nil
}

