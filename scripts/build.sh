#!/usr/bin/env bash

GEN_COUNT=$(pokego named generation --count)
SPECIES_COUNT=$(pokego named pokemon-species --count)
declare -A GEN_DOCS

# curl https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/dream-world/155.svg -o website/static/img/logo.svg

for ((i = 1; i <= $GEN_COUNT; i++)); do
    mkdir -p website/docs/gen$i
    mkdir -p website/static/sprites/pokemon/official-artwork
    mkdir -p website/static/sprites/pokemon/front-default
    gen=$(pokego named generation --raw | jq -r ".results[$i-1].name")
    # echo $gen

    GEN_DOCS[$gen]=gen$i
    # echo $GEN_DOCS
done

for ((i = 1; i <= 3; i++)); do
    pokemon_gen=$(pokego species $i -g | jq -r '.name')
    pokemon_name=$(pokego pokemon $i --name)
    if [[ ! -e website/docs/${GEN_DOCS[$pokemon_gen]}/$pokemon_name.md.md ]]; then
        pokego render --id $i -i tmpl/pokemon.md.tmpl -o website/docs/${GEN_DOCS[$pokemon_gen]}/$pokemon_name.md
    fi

    if [[ ! -e website/static/sprites/pokemon/official-artwork/$pokemon_name.png ]]; then
        artwork_link=$(pokego pokemon $i --official-artwork)
        curl $artwork_link -o website/static/sprites/pokemon/official-artwork/$pokemon_name.png
    fi

    if [[ ! -e website/static/sprites/pokemon/front-default/$pokemon_name.png ]]; then
        artwork_link=$(pokego pokemon $i --front-default)
        curl $artwork_link -o website/static/sprites/pokemon/front-default/$pokemon_name.png
    fi
done

# first get gen count, and make doc folders
# then for each gen number, get pokemon x-y
# for each pokemon p in x-y, generate doc
