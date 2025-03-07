package main

import (
	"fmt"

	// "github.com/maquentr/REPL/internal/pokeapi"
)

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("Pokemon in Pokedex:")
	for _, poke := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", poke.Name )
	}
	return nil
}
