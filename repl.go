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
	callback    func(*config, ...string) error
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
			description: "List next page of location areas",
			callback:    CallbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous page of location areas",
			callback:    CallbackMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "List the Pokemon in a location area",
			callback:    CallbackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Try to catch the darn thing & add to your Pokedex",
			callback:    CallbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "View information about a caught Pokemon",
			callback:    CallbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all Pokemon in your Pokedex",
			callback:    CallbackPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit PokeDex",
			callback:    CallbackExit,
		},
	}
}
