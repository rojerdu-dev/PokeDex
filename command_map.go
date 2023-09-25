package main

import (
	"fmt"
	"log"
)

func CallbackMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.GetLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	return nil
}
