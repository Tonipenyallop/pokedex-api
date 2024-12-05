package pokemonServiceHelper

import (
	"fmt"

	pokemonRepository "github.com/Tonipenyallop/pokedex-api/repository"
	"github.com/patrickmn/go-cache"
)

const POKEMON_CACHE_KEY = "pokemons"


func GetPokemonsFromCacheByGen(genId string, pokemonCache *cache.Cache) []pokemonRepository.TmpPokemon {
	genCacheKey := GetGenCacheKey(genId)

	pokemonInCache, found := pokemonCache.Get(genCacheKey)
	fmt.Println("found",found)
	if found {
		if cachedPokemon, ok := pokemonInCache.([]pokemonRepository.TmpPokemon); ok {
			return cachedPokemon
		}
	}
	return nil
}



func GetGenCacheKey(genId string)string{
	return fmt.Sprintf("%s%s%s", POKEMON_CACHE_KEY, '_', genId)
}

func GetGenIdByPokemonId(pokemonId int) string{
	
	if 1 <= pokemonId && pokemonId <= 151 {

		
		return "first"

	}
	if 152 <= pokemonId && pokemonId <= 251 {

		
		return "second"

	}
	if 252 <= pokemonId && pokemonId <= 386 {

		
		return "third"

	}
	if 387 <= pokemonId && pokemonId <= 493 {

		
		return "fourth"

	}
	if 494 <= pokemonId && pokemonId <= 649 {

		
		return "fifth"

	}
	if 650 <= pokemonId && pokemonId <= 721 {

		
		return "sixth"

	}
	if 722 <= pokemonId && pokemonId <= 809 {

		
		return "seventh"

	}
	if 810 <= pokemonId && pokemonId <= 905 {

		
		return "eighth"

	}
	if 906 <= pokemonId && pokemonId <= 1025 {

		
		return "ninth"

	}
	
	return 	""
	
	
}