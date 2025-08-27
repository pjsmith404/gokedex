package main

import (
	"fmt"
)

func commandInspect(conf *config, subCommand string) error {
	if subCommand == "" {
		return fmt.Errorf("Pick a pokemon to inspect. eg: inspect pikachu")
	}

	pokemon := conf.pokedex[subCommand]

	if pokemon.ID == 0 {
		return fmt.Errorf("%v has not been caught!", subCommand)
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, pType := range pokemon.Types {
		fmt.Println(" -", pType.Type.Name)
	}

	return nil
}
