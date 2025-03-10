package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name		string
	description	string
	callback	func(*config, ...string) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" >")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Prints the help menu",
			callback: callbackHelp,
		},
		"map": {
			name: "map",
			description: "Lists of location areas",
			callback: callbackMap,
		},
		"mapb": {
			name: "mapb",
			description: "Lists of previous location areas",
			callback: callbackMapb,
		},
		"explore": {
			name: "explore {location_area}",
			description: "Lists the pokemon in a location area",
			callback: callbackExplore,
		},
		"catch": {
			name: "catch {pokemon_name}",
			description: "attempt to catch a pokemon and add it to your pokedex",
			callback: callbackCatch,
		},
		"inspect": {
			name: "inspect {pokemon_name}",
			description: "view information about caught pokemon",
			callback: callbackInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "view all the pokemon in your pokedex",
			callback: callbackPokedex,
		},
		"exit": {
			name: "exit",
			description: "exits prompt",
			callback: callbackExit,
		},
	}
}


func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
