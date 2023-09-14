#!/usr/bin/env bash

if ! command -v pokego &>/dev/null; then
    curl -sL https://github.com/jeremyhager/pokego/releases/latest/download/mac.gz | tar xz
    cp pokego /usr/local/bin
fi

GEN_COUNT=$(pokego named generation --count)
SPECIES_COUNT=$(pokego named pokemon-species --count)
declare -A GEN_DOCS

for ((i = 1; i <= $GEN_COUNT; i++)); do
    mkdir -p docs/gen$i static/sprites/pokemon/official-artwork static/sprites/pokemon/front-default
    gen=$(pokego named generation --raw | jq -r ".results[$i-1].name")

    GEN_DOCS[$gen]=gen$i
done

for ((i = 1; i <= 3; i++)); do
    pokemon_gen=$(pokego species $i -g | jq -r '.name')
    pokemon_name=$(pokego pokemon $i --name)
    if [[ ! -e docs/${GEN_DOCS[$pokemon_gen]}/$pokemon_name.md.md ]]; then
        pokego render --id $i -i tmpl/pokemon.md.tmpl -o docs/${GEN_DOCS[$pokemon_gen]}/$pokemon_name.md
    fi

    if [[ ! -e static/sprites/pokemon/official-artwork/$pokemon_name.png ]]; then
        artwork_link=$(pokego pokemon $i --official-artwork)
        curl $artwork_link -o static/sprites/pokemon/official-artwork/$pokemon_name.png
    fi

    if [[ ! -e static/sprites/pokemon/front-default/$pokemon_name.png ]]; then
        artwork_link=$(pokego pokemon $i --front-default)
        curl $artwork_link -o static/sprites/pokemon/front-default/$pokemon_name.png
    fi
done
