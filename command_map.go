package main

import (
	"fmt"
	"github.com/rojerdu-dev/PokeDex/internal/pokeapi"
	"log"
)

func CallbackMap() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.GetLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
