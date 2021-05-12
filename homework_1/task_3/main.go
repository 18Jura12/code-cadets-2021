package main

import (
	"fmt"
	"task3/pokemon_encounters"
)

func main() {

	err := pokemonencounters.GetAreas()
	if err != nil {
		fmt.Println(err)
	}

}