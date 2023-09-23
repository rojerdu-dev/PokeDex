package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	startRepl()
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
		command := cleaned[0]

		switch command {
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("invalid command")
		}

		//fmt.Println("right back at you:", cleaned)

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
