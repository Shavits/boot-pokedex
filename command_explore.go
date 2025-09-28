package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, param string) error {

	if param == ""{
		return errors.New("you must provide a location name")
	}
	locationsResp, err := cfg.pokeapiClient.GetLocationData(param)
	if err != nil {
		return err
	}

	for _, enc := range locationsResp.PokemonEncounters {
		fmt.Println(enc.Pokemon.Name)
	}
	return nil

}