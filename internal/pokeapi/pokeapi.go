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

type ResLocationAreaDetail struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues *[]any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

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
