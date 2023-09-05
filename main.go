package main

import (
	"log"
	"os"
	"text/template"

	"github.com/jeremyhager/pokego/pkg/pokedex"
)

func main() {

	info, err := pokedex.Init("155")
	if err != nil {
		log.Fatal(err)
	}

	// type pokemonInfo struct {
	// 	Pokemon          pokemon.Pokemon
	// 	Species          pokemonspecies.PokemonSpecies
	// 	Evolution        evolutionchains.EvolutionChain
	// 	EvolutionSpecies []pokemonspecies.PokemonSpecies
	// 	EvolutionPokemon []pokemon.Pokemon
	// 	// EvolutionOptions []evolutionchains.ChainLink
	// }

	// // 	<p>
	// // 	{{ $details := (index $.EvolutionOptions $i).EvolutionDetails -}}
	// // 	  {{- range $detail := $details -}}
	// // 		{{- if . -}}
	// // 		  {{- . -}}
	// // 		{{- end -}}
	// // 	  {{- end }}
	// //   </p>

	// poke, err := pokemon.Get("ditto")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// speciesInfo, err := pokemonspecies.Get(poke.Species.Name)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// endpoint, err := url.Parse(speciesInfo.EvolutionChain.Url)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// chainId := strings.Split(endpoint.Path, "/")[4]

	// evolutions, err := evolutionchains.Get(chainId)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// pokedex := pokemonInfo{
	// 	Pokemon: poke,
	// 	Species: speciesInfo,
	// }

	// baseSpecies, err := pokedex.Species.GetBaseSpecies(&evolutions.Chain)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pokedex.EvolutionSpecies = append(pokedex.EvolutionSpecies, baseSpecies)

	// if len(evolutions.Chain.EvolvesTo) > 0 {

	// 	evolutionChainPokemon, err := pokedex.Species.FlattenEvolutions(&evolutions.Chain)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	pokedex.Evolution = evolutions
	// 	pokedex.EvolutionSpecies = append(pokedex.EvolutionSpecies, evolutionChainPokemon...)
	// 	for _, species := range evolutionChainPokemon {
	// 		for _, defaultVarity := range species.Varieties {
	// 			if defaultVarity.IsDefault {
	// 				evolutionTo, err := pokemon.Get(defaultVarity.Pokemon.Name)
	// 				if err != nil {
	// 					log.Fatal(err)
	// 				}
	// 				pokedex.EvolutionPokemon = append(pokedex.EvolutionPokemon, evolutionTo)
	// 			}
	// 		}
	// 	}

	// } else {
	// 	// do nothing
	// }
	tmplFile, err := template.ParseFiles("website/docs/gen2/pokemon.md.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	err = tmplFile.Execute(os.Stdout, info)

	if err != nil {
		log.Fatal(err)
	}

}
