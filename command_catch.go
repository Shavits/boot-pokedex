package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, param string) error {

	pokemonResp, err := cfg.pokeapiClient.GetPokemonData(param)
	if err != nil {
		return err
	}

	
	baseExperience := pokemonResp.BaseExperience
	catchAttempt := rand.Intn(256)
	
	
	fmt.Printf("Throwing a Pokeball at %s...\n", param)
	if baseExperience <= catchAttempt{
		fmt.Printf("%s was caught!\n", param)
		cfg.pokeDex[param] = pokemonResp
	}else{
		fmt.Printf("%s escaped!\n", param)
	}
	return nil

}

