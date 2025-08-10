package pokeapi

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
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

func (c *Client) GetLocationArea(pageUrl *string) (ResLocationArea, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	data, found := c.cache.Get(url)
	fmt.Println(data, found)

	if !found {
		fmt.Println("here")
		res, err := http.Get(url)
		if err != nil {
			return ResLocationArea{}, err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			fmt.Errorf(
				"Response failed with status code: %d and\nbody: %s\n",
				res.StatusCode,
				res.Body,
			)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return ResLocationArea{}, err
		}

		c.cache.Add(url, data)
	}

	var locationArea ResLocationArea
	err := json.Unmarshal(data, &locationArea)
	if err != nil {
		return ResLocationArea{}, err
	}

	return locationArea, nil
}

