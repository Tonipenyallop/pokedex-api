package pokemonService

import (
	"fmt"
	"sort"

	pokemonRepository "github.com/Tonipenyallop/pokedex-api/repository"
)



func GetAllPokemons() ([]pokemonRepository.TmpPokemon, error) {

	pokemons, err := pokemonRepository.GetAllPokemons()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all pokemons:",err)
	}


	// Sort by ID
	sort.Slice(pokemons, func(i, j int) bool {
		return pokemons[i].ID < pokemons[j].ID
	})

	return pokemons , nil


}

func GetPokemonsByGen(genId string) ([]pokemonRepository.TmpPokemon, error) {
	pokemons, err := pokemonRepository.GetPokemonsByGen(genId)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch all pokemons:",err)
	}


	// Sort by ID
	sort.Slice(pokemons, func(i, j int) bool {
		return pokemons[i].ID < pokemons[j].ID
	})

	return pokemons , nil
}