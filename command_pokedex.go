package main

import (
	"fmt"
)

func CallbackPokedex(cfg *config, args ...string) error {
	fmt.Println("Pokemon in PokeDex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
