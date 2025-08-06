package main

import (
	"fmt"
	"github.com/pjsmith404/gokedex/internal/pokeapi"
)

func commandMap(conf *config) error {
	locationArea, err := pokeapi.GetLocationArea(conf.next)
	if err != nil {
		return err
	}

	for _, location := range locationArea.Results {
		fmt.Println(location.Name)
	}

	conf.next = locationArea.Next
	conf.previous = locationArea.Previous

	return nil
}

func commandMapBack(conf *config) error {
	if conf.previous == nil {
		return fmt.Errorf("You're on the first page")
	}

	locationArea, err := pokeapi.GetLocationArea(conf.previous)
	if err != nil {
		return err
	}

	for _, location := range locationArea.Results {
		fmt.Println(location.Name)
	}

	conf.next = locationArea.Next
	conf.previous = locationArea.Previous

	return nil
}
