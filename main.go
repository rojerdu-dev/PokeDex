package main

import (
	"github.com/rojerdu-dev/PokeDex/internal/pokeapi"
	"time"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}

	StartRepl(&cfg)
}
