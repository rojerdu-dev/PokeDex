package main

import (
	"fmt"
)

func CallbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to PokeDex Help Menu!")
	fmt.Println("Available commands:")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println("")
	return nil
}
