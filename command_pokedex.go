package main

import (
	"fmt"
)

func commandPokedex(conf *config, subCommand string) error {
	fmt.Println("Your Pokedex:")
	for pokemon, _ := range conf.pokedex {
		fmt.Println(" -", pokemon)
	}

	return nil
}
