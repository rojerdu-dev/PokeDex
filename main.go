package main

import (
	"bufio"
	"fmt"
	"github.com/rojerdu-dev/PokeDex/internal/pokeapi"
	"log"
	"os"
	"strings"
)

func main() {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.GetLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	//startRepl()
}

func startRepl() {
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
		command.callback()

		if inputText == "exit" {
			fmt.Println("sayonara")
			break
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints this help menu",
			callback:    CallbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit PokeDex",
			callback:    CallbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
