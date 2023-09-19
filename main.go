package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your name: ")

	for scanner.Scan() {
		inputText := scanner.Text()
		if inputText == "exit" {
			fmt.Println("Exiting...")
			break
		}
		fmt.Println("You entered:", inputText)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
