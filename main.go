package main

import "github.com/rojerdu-dev/PokeDex/internal/pokeapi"

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	StartRepl(&cfg)
}
