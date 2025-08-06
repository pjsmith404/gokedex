package main

import (
	"fmt"
	"encoding/json"
	"github.com/pjsmith404/gokedex/internal/pokeapi"
)

func commandMap(conf *config) error {
	var url string
	if conf.next == "" {
		url = baseUrl + "location-area/"
	} else {
		url = conf.next
	}

	res := pokeapi.Get(url)
	locationArea := LocationArea{}
	err := json.Unmarshal(res, &locationArea)
	if err != nil {
		fmt.Println(err)
	}

	for _,location := range locationArea.Results {
		fmt.Println(location.Name)
	}

	conf.next = locationArea.Next
	conf.previous = locationArea.Previous

	return nil
}

func commandMapBack(conf *config) error {
	var url string
	if conf.previous == "" {
		url = baseUrl + "location-area/"
	} else {
		url = conf.previous
	}

	res := pokeapi.Get(url)
	locationArea := LocationArea{}
	err := json.Unmarshal(res, &locationArea)
	if err != nil {
		fmt.Println(err)
	}

	for _,location := range locationArea.Results {
		fmt.Println(location.Name)
	}

	conf.next = locationArea.Next
	conf.previous = locationArea.Previous

	return nil
}

