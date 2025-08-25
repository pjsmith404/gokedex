package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(conf *config, subCommand string) error {
	if subCommand == "" {
		return fmt.Errorf("Pick a pokemon to catch! eg: catch pikachu")
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", subCommand)

	pokemon, err := conf.pokeapiClient.GetPokemon(subCommand)
	if err != nil {
		return err
	}

	randInt := rand.Intn(1000)
	if randInt > pokemon.BaseExperience {
		fmt.Printf("%v was caught!\n", subCommand)
		conf.pokedex[subCommand] = pokemon
	} else {
		fmt.Printf("%v escaped!\n", subCommand)
	}

	return nil
}
