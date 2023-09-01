package main

import (
	"log"
	"os"
	"text/template"

	"github.com/jeremyhager/pokeapi/pokemon"
	"github.com/jeremyhager/pokeapi/pokemonspecies"
)

func main() {
	type pokemonInfo struct {
		Pokemon pokemon.Pokemon
		Species pokemonspecies.PokemonSpecies
	}

	poke, err := pokemon.Get("bulbasaur")
	if err != nil {
		log.Fatal(err)
	}

	speciesInfo, err := pokemonspecies.Get(poke.Species.Name)
	if err != nil {
		log.Fatal(err)
	}

	pokedex := pokemonInfo{
		Pokemon: poke,
		Species: speciesInfo,
	}

	tmplFile, err := template.ParseFiles("website/docs/gen2/pokemon.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	err = tmplFile.Execute(os.Stdout, pokedex)

	// err = tmpl.Execute(os.Stdout, poke)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("pokemon: %+v", poke.Abilities)
}
