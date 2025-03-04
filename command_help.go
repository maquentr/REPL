package main

import "fmt"

func callbackHelp(cfg *config) error {
	fmt.Println("Welcone to the repl help menu")
	fmt.Println("Available commands:")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}