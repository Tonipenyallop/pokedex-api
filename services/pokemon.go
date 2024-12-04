package pokemonService

import (
	"fmt"
	"sort"
	"time"

	pokemonRepository "github.com/Tonipenyallop/pokedex-api/repository"
	"github.com/patrickmn/go-cache"
)

const POKEMON_CACHE_KEY = "pokemons"

var pokemonCache = cache.New(24*time.Hour, 24*time.Hour) // Shared cache instance

func GetAllPokemons() ([]pokemonRepository.TmpPokemon, error) {
	// Check cache
	cachedPokemon := getAllPokemonsFromCache()
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
	// Check cache for all pokemons
	cachedPokemon := getAllPokemonsFromCache()
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

	// Check cache for generation pokemons
	cachedGenPokemons := getPokemonsFromCacheByGen(genId)
	if cachedGenPokemons != nil {
		return cachedGenPokemons, nil
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
	genCacheKey := fmt.Sprintf("%s%s%s", POKEMON_CACHE_KEY, '_', genId)
	pokemonCache.Set(genCacheKey, pokemons, 24*time.Hour)

	return pokemons, nil
}

func getAllPokemonsFromCache() []pokemonRepository.TmpPokemon {
	pokemonInCache, found := pokemonCache.Get(POKEMON_CACHE_KEY)
	if found {
		if cachedPokemon, ok := pokemonInCache.([]pokemonRepository.TmpPokemon); ok {
			return cachedPokemon
		}
	}
	return nil
}

func getPokemonsFromCacheByGen(genId string) []pokemonRepository.TmpPokemon {
	genCacheKey := fmt.Sprintf("%s%s%s", POKEMON_CACHE_KEY, '_', genId)

	pokemonInCache, found := pokemonCache.Get(genCacheKey)
	if found {
		if cachedPokemon, ok := pokemonInCache.([]pokemonRepository.TmpPokemon); ok {
			return cachedPokemon
		}
	}
	return nil
}
