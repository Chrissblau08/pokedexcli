package main

import (
	"fmt"
	"os"
	"sort"
)

// Signatur aller Command-Funktionen
type CommandFunc func() error

// CLI Command struct
type cliCommand struct {
	name        string
	description string
	callback    CommandFunc
}

// Command: Help
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	// Sortierte Ausgabe
	names := make([]string, 0, len(commands))
	for name := range commands {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		cmd := commands[name]
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

// Command: Exit
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// Zuerst die Map deklarieren (leer)
var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
