package pokemonService

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	pokemonServiceHelper "github.com/Tonipenyallop/pokedex-api/helpers"
	pokemonRepository "github.com/Tonipenyallop/pokedex-api/repository"
	"github.com/Tonipenyallop/pokedex-api/types"
	"github.com/patrickmn/go-cache"
)

const POKEMON_CACHE_KEY = "pokemons"

var pokemonCache = cache.New(24*time.Hour, 24*time.Hour) // Shared cache instance

func GetAllPokemons() ([]pokemonRepository.TmpPokemon, error) {
	// Check cache
	cachedPokemon := pokemonServiceHelper.GetAllPokemonsFromCache(pokemonCache)
	if cachedPokemon != nil {
		return cachedPokemon, nil
	}

	// Fetch from repository
	pokemons, err := pokemonRepository.GetAllPokemons()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all pokemons: %w", err)
	}

	// Sort and cache
	sort.Slice(pokemons, func(i, j int) bool {
		return pokemons[i].ID < pokemons[j].ID
	})

	// store cache for All pokemons
	pokemonCache.Set(POKEMON_CACHE_KEY, pokemons, 24*time.Hour)

	return pokemons, nil
}

func GetPokemonsByGen(genId string) ([]pokemonRepository.TmpPokemon, error) {
	// Check cache for generation pokemons
	cachedGenPokemons := pokemonServiceHelper.GetPokemonsFromCacheByGen(genId, pokemonCache)

	if cachedGenPokemons != nil {
		return cachedGenPokemons, nil
	}

	// Check cache for all pokemons
	cachedPokemon := pokemonServiceHelper.GetAllPokemonsFromCache(pokemonCache)
	var pokemonArr []pokemonRepository.TmpPokemon
	if cachedPokemon != nil {

		for _, pokemon := range cachedPokemon {
			if pokemon.Generation == genId {
				// pokemon already sorted when storing cache
				pokemonArr = append(pokemonArr, pokemon)
			}
		}
		return pokemonArr, nil
	}

	// Fetch from repository
	pokemons, err := pokemonRepository.GetPokemonsByGen(genId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pokemons by generation: %w", err)
	}

	// Sort and cache
	sort.Slice(pokemons, func(i, j int) bool {
		return pokemons[i].ID < pokemons[j].ID
	})

	// store cache for generation
	genCacheKey := pokemonServiceHelper.GetGenCacheKey(genId)
	pokemonCache.Set(genCacheKey, pokemons, 24*time.Hour)

	return pokemons, nil
}

// func GetEvolutionChainById(pokemonId string) (*types.EvolutionChain, error) {

// 	evolutionChain, err := pokemonRepository.GetEvolutionChainById(pokemonId)
// 	if err != nil {
// 		return nil, fmt.Errorf("Failed to get evolution chain")
// 	}

// 	return evolutionChain, nil

// }

// get flavor text and evolution_chain

func GetPokemonFlavorTextAndEvolutionChain(pokemonId string) (*types.SpeciesInfo, error) {
	flavorText, err := pokemonRepository.GetPokemonFlavorTextAndEvolutionChain(pokemonId)
	if err != nil {
		return nil, fmt.Errorf("Failed to get flavor text")
	}

	tmp, err := pokemonRepository.GetEvolutionChain(flavorText.EvolutionChain.URL)

	if err != nil {
		return nil, fmt.Errorf("Failed to fetch evolution chain detail")
	}
	fmt.Println("tmp", tmp)

	// evolutionChain, err := pokemonRepository.GetEvolutionChainById(pokemonId)
	// if err != nil {
	// 	return nil, fmt.Errorf("Failed to get evolution chain")
	// }

	// [types.SpeciesInfo, ]

	return flavorText, nil

}

func GetPokemonDetail(pokemonId string) (*pokemonRepository.TmpPokemon, error) {
	convertedPokemonId, err := strconv.Atoi(pokemonId)

	// check gen cache first
	genId := pokemonServiceHelper.GetGenIdByPokemonId(convertedPokemonId)

	// get pokemon from gen cache
	genPokemonsInCache := pokemonServiceHelper.GetPokemonsFromCacheByGen(genId, pokemonCache)

	if genPokemonsInCache != nil {
		for _, pokemonInCache := range genPokemonsInCache {
			if pokemonInCache.ID == convertedPokemonId {
				pokemon := pokemonRepository.TmpPokemon(pokemonInCache)
				fmt.Println("pokemon found it!", pokemon)
				return &pokemon, nil
			}
		}

	}

	// check all gen cache
	allPokemonsInCache := pokemonServiceHelper.GetAllPokemonsFromCache(pokemonCache)

	
	if allPokemonsInCache != nil {
		for _, pokemonInCache := range allPokemonsInCache {
			if pokemonInCache.ID == convertedPokemonId {
				pokemon := pokemonRepository.TmpPokemon(pokemonInCache)
				return &pokemon, nil
			}
		}
	}

	pokemon, err := pokemonRepository.GetPokemonDetail(pokemonId)

	if err != nil {
		return nil, fmt.Errorf("Failed to get pokemon detail", err)
	}

	return pokemon, nil

}
