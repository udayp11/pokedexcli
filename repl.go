package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func enterRepl() {

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
			err := command.callback()
			if err != nil {
				fmt.Errorf("error: %w", err)
			}

		}

	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	finalStrings := strings.Fields(lowerText)
	return finalStrings

}
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp() error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage")
	fmt.Println()
	for _, cmds := range inputCmd() {
		fmt.Printf("%v: %v\n", cmds.name, cmds.description)
	}
	return nil

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
	}
}
