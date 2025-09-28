package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, param string) error {
	if param == ""{
		return errors.New("you must provide a pokemon name")
	}

	poekmon, caught := cfg.pokeDex[param]
	if !caught{
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Println("Name:", poekmon.Name)
	fmt.Println("Height:", poekmon.Height)
	fmt.Println("Weight:", poekmon.Weight)
	fmt.Println("Stats:")
	for _, stat := range(poekmon.Stats){
		fmt.Printf("	-%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range(poekmon.Types){
		fmt.Printf("	- %v\n", t.Type.Name)
	}
	


	return nil
}
