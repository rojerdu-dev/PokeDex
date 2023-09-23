package main

import (
	"fmt"
)

func CallbackHelp() {
	fmt.Println("Welcome to PokeDex Help Menu!")
	fmt.Println("Available commands:")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("- %s: %s", cmd.name, cmd.description)
	}

	fmt.Println("")
}
