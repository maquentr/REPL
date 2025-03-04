package main

import (
	// "fmt"
	// "log"

	"github.com/maquentr/REPL/internal/pokeapi"
)

type config struct {
	pokeapiClient		pokeapi.Client
	nextLocationAreaURL	*string
	prevLocationAreaURL	*string
}

func main() {
	cfg := config {
		pokeapiClient: pokeapi.NewClient(),
	}
	
	startRepl(&cfg)
}
