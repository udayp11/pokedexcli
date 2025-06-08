package main

import (
	"fmt"
)

func commandHelp(cfg *config, area string) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage")
	fmt.Println()
	for _, cmds := range inputCmd() {
		fmt.Printf("%v: %v\n", cmds.name, cmds.description)
	}
	return nil
}
