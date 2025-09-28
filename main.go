package main

import (
	"time"

	"github.com/shavits/boot-pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokeDex: map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)
}
