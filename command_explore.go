package main

import (
	"fmt"
)

func commandExplore(cfg *config, area string) error {

	pokemons, err := cfg.pokeapiClient.GetPokemon(area)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", pokemons.Name)
	fmt.Println("Found Pokemon: ")

	for _, result := range pokemons.PokemonEncounters {
		fmt.Println(result.Pokemon.Name)
	}

	return nil

}
