package main

import (
	"bufio"
	"fmt"
	"os"
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

		fmt.Println("right back at you:", inputText)

		if inputText == "exit" {
			fmt.Println("sayonara")
			break
		}
	}

}
