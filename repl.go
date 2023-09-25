package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" >")
		scanner.Scan()
		inputText := scanner.Text()

		cleaned := cleanInput(inputText)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		command.callback(cfg)

		if inputText == "exit" {
			fmt.Println("sayonara")
			break
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints this help menu",
			callback:    CallbackHelp,
		},
		"map": {
			name:        "map",
			description: "List some location areas",
			callback:    CallbackMap,
		},
		"exit": {
			name:        "exit",
			description: "Exit PokeDex",
			callback:    CallbackExit,
		},
	}
}
