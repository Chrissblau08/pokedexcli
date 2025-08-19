package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		words := cleanInput(input)

		if len(words) == 0 {
			continue
		}

		cmdName := words[0]

		// Prüfen, ob Command existiert
		if cmd, ok := commands[cmdName]; ok {
			if err := cmd.callback(); err != nil {
				fmt.Printf("Error executing command %s: %v\n", cmdName, err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", cmdName)
		}
	}
}
