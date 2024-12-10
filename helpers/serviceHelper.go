package pokemonServiceHelper

import (
	"fmt"
	"strings"
	"unicode"

	pokemonRepository "github.com/Tonipenyallop/pokedex-api/repository"
	"github.com/Tonipenyallop/pokedex-api/types"
	"github.com/patrickmn/go-cache"
)

const POKEMON_CACHE_KEY = "pokemons"


func GetPokemonsFromCacheByGen(genId string, pokemonCache *cache.Cache) []pokemonRepository.TmpPokemon {
	genCacheKey := GetGenCacheKey(genId)

	pokemonInCache, found := pokemonCache.Get(genCacheKey)
	if found {
		if cachedPokemon, ok := pokemonInCache.([]pokemonRepository.TmpPokemon); ok {
			return cachedPokemon
		}
	}
	return nil
}


func GetAllPokemonsFromCache(pokemonCache*cache.Cache) []pokemonRepository.TmpPokemon {
	pokemonInCache, found := pokemonCache.Get(POKEMON_CACHE_KEY)
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



func GetEvolutionChainPokemonNames(evolutionChain *types.EvolutionChain)[]string{

	var output []string
	output = append(output,evolutionChain.Chain.Species.Name) 

	

	var queue []types.Chain
	
	for _, pokemon := range evolutionChain.Chain.EvolvesTo{
		output = append(output, pokemon.Species.Name)
		queue = append(queue, pokemon.EvolvesTo...)
	}
	// for i := 0 ; i < 2; i++ {

	// }



	return  output
}


func HelperDescription(s string)[]types.YoutubeMusic{
	

	output := make([]types.YoutubeMusic,106)
	// for array index
	counter := 0
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// Split by whitespace
		fields := strings.Fields(line)

		// We want lines starting with a number and ending with a time, e.g.:
		// "1 タイトルデモ ～ホウエン地方の旅立ち～ 00:12"
		// Check if first field is a number and last field looks like a time.
		if len(fields) < 3 {
			continue
		}

		// Check if the first field is all digits
		if !isAllDigits(fields[0]) {
			continue
		}

		// Time is last field
		time := fields[len(fields)-1]
		// The title is everything between the first and last
		title := strings.Join(fields[1:len(fields)-1], " ")

		songName := fmt.Sprintf("%s %s",fields[0],title)
		output[counter].Name = songName
		output[counter].StartTime = time
		counter ++
	}
	return output
}

func isAllDigits(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}