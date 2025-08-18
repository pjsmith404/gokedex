package main

import (
	"fmt"
)

func commandMap(conf *config, subCommand string) error {
	locationArea, err := conf.pokeapiClient.GetLocationArea(conf.next)
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

func commandMapBack(conf *config, subCommand string) error {
	if conf.previous == nil {
		return fmt.Errorf("You're on the first page")
	}

	locationArea, err := conf.pokeapiClient.GetLocationArea(conf.previous)
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
