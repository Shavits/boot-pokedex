package main

import (
	"fmt"
)

func commandExplore(cfg *config, param string) error {

	locationsResp, err := cfg.pokeapiClient.GetLocationData(param)
	if err != nil {
		return err
	}

	for _, enc := range locationsResp.PokemonEncounters {
		fmt.Println(enc.Pokemon.Name)
	}
	return nil

}