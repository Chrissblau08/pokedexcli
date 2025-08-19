package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
)

var pokedex = make(map[string]PokemonCatch)

// Signatur aller Command-Funktionen
type CommandFunc func(cfg *Config, args ...string) error

// CLI Command struct
type cliCommand struct {
	name        string
	description string
	callback    CommandFunc
}

type Config struct {
	Next     string
	Previous string
}

// Command: Map
func commandMap(cfg *Config, args ...string) error {
	url := cfg.Next
	locations, next, previous, err := getLocationsArea(url)
	if err != nil {
		return err
	}

	for _, loc := range locations {
		fmt.Println(loc)
	}

	cfg.Next = next
	cfg.Previous = previous

	return nil
}

// Command: Map
func commandMapb(cfg *Config, args ...string) error {
	url := cfg.Previous
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, next, previous, err := getLocationsArea(url)
	if err != nil {
		return err
	}

	for _, loc := range locations {
		fmt.Println(loc)
	}

	// Config aktualisieren
	cfg.Next = next
	cfg.Previous = previous

	return nil
}

func commandExplore(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide an area name, e.g. explore canalave-city")
	}

	areaName := args[0]

	fmt.Println("Exploring " + areaName + "...")

	pokemonEncounters, err := getPokemonEncountersbyLocationArea(areaName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonEncounters {
		fmt.Println(" - " + pokemon)
	}
	return nil
}

// Helper Function for Catch - Command
func catchPokemon(baseExp int) bool {
	// Berechne Fangchance: z.B. maxChance = 90%
	// je höher baseExp, desto niedriger die Chance
	maxChance := 90
	chance := maxChance - baseExp/2
	if chance < 5 {
		chance = 5 // Mindestens 5% Chance
	}

	roll := rand.Intn(100) // 0-99
	return roll < chance
}

// Command: Catch
func commandCatch(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide an pokemon name, e.g. pikachu")
	}

	pokemonName := args[0]
	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")

	pokemonObject, err := getPokemonByNameInfo(pokemonName)
	if err != nil {
		return err
	}

	if catchPokemon(pokemonObject.BaseExperience) {
		pokedex[pokemonName] = pokemonObject
		fmt.Println(pokemonName + " was caught!")
	} else {
		fmt.Println(pokemonName + " escaped!")
	}

	return nil
}

// Command: Inspect
func commandInspect(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide an pokemon name that you catched, e.g. pikachu")
	}

	pokemonName := args[0]
	pokemonObject, ok := pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("pokemon %s not found in your pokedex", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemonObject.Name)
	fmt.Printf("Height: %d\n", pokemonObject.Height)
	fmt.Printf("Weight: %d\n", pokemonObject.Weight)

	fmt.Println("Stats:")
	for _, s := range pokemonObject.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemonObject.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}

// Command: Help
func commandHelp(cfg *Config, args ...string) error {
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
func commandExit(cfg *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandPokedex(cfg *Config, args ...string) error {
	keys := make([]string, 0, len(pokedex))
	for k := range pokedex {
		keys = append(keys, k)
	}

	if len(keys) == 0 {
		return fmt.Errorf("your Pokedex is empty :C")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemonName := range keys {
		fmt.Println(" - " + pokemonName)
	}
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
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapd (map back)",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List of all the Pokémon located there",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catching Pokemon adds them to the user's Pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "It takes the name of a Pokemon and prints the name, height, weight, stats and type(s) of the Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Prints a list of all the names of the Pokemon the user has caught",
			callback:    commandPokedex,
		},
	}
}
