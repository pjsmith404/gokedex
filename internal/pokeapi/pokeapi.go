package pokeapi

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
)

const baseUrl = "https://pokeapi.co/api/v2"

func (c *Client) GetLocationArea(pageUrl *string) (ResLocationArea, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	data, ok := c.cache.Get(url)

	if !ok {
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

func (c *Client) GetLocationAreaDetail(id string) (ResLocationAreaDetail, error) {
	url := baseUrl + "/location-area/" + id

	data, ok := c.cache.Get(url)

	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return ResLocationAreaDetail{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return ResLocationAreaDetail{}, err
		}

		if res.StatusCode > 299 {
			return ResLocationAreaDetail{}, fmt.Errorf(
				"Response failed with status code: %v\nBody:\n%v",
				res.StatusCode,
				string(data),
			)
		}

		c.cache.Add(url, data)
	}

	var locationAreaDetail ResLocationAreaDetail
	err := json.Unmarshal(data, &locationAreaDetail)
	if err != nil {
		return ResLocationAreaDetail{}, fmt.Errorf("Error unmarshalling data: %w", err)
	}

	return locationAreaDetail, nil
}
