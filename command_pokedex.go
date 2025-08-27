package main

import (
	"fmt"
)

func commandPokedex(conf *config, subCommand string) error {
	fmt.Println("Your Pokedex:")
	for _, p := range conf.pokedex {
		fmt.Println(" -", p.Name)
	}

	return nil
}
