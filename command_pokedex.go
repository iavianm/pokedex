package main

import (
	"fmt"
)

func commandPokedex(cfg *config, _ ...string) error {

	fmt.Println("Your Pokedex:")
	for key, _ := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n", key)
	}

	return nil
}
