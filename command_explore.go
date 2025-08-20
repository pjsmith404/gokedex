package main

import (
	"fmt"
)

func commandExplore(conf *config, subCommand string) error {
	if subCommand == "" {
		return fmt.Errorf("Provide an area to explore. ie: explore canalave-city-area")
	}
	fmt.Println("Explore!")

	return nil
}
