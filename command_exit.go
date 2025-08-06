package main

import (
	"os"
	"fmt"
)

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return fmt.Errorf("Failed to exit program")
}

