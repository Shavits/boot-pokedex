package main

import (
	"fmt"
)

func commandPokedex(cfg *config, param string) error {

	if len(cfg.pokeDex) == 0{
		fmt.Println("You have not caught any pokemon yet...")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for k := range(cfg.pokeDex){
		fmt.Printf("	- %v\n", k)
	}
	


	return nil
}
