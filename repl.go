package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/udayp11/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}
type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
	caughtPokemon   map[string]pokeapi.Pokemon
}

func enterRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		userInput := scanner.Text()
		userCmd := cleanInput(userInput)
		if len(userCmd) == 0 {
			continue
		}
		command, exists := inputCmd()[userCmd[0]]
		if !exists {
			fmt.Println("Unknown command")
		} else {
			area := ""
			if len(userCmd) > 1 {
				area = userCmd[1]
			}
			err := command.callback(cfg, area)
			if err != nil {
				fmt.Println(err)
			}

		}

	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	finalStrings := strings.Fields(lowerText)
	return finalStrings

}

func inputCmd() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 Locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 Locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "list of all the Pokémon Present in a area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch the Pokemon by name",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Get details of the catched pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Get List of all Pokemons caught",
			callback:    commandPokedex,
		},
	}
}
