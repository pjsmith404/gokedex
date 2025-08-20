package main

import (
	"fmt"
)

func commandExplore(conf *config, subCommand string) error {
	if subCommand == "" {
		return fmt.Errorf("Provide an area to explore. eg: explore canalave-city-area")
	}

	fmt.Printf("Exploring %v...\n", subCommand)

	locationAreaDetail, err := conf.pokeapiClient.GetLocationAreaDetail(subCommand)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationAreaDetail.PokemonEncounters {
		fmt.Println("-", pokemon.Pokemon.Name)
	}

	return nil
}
