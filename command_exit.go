package main

import (
	"fmt"
	"os"
)

func commandExit(conf *config, subCommand string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return fmt.Errorf("Failed to exit program")
}
