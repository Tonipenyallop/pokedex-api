package pokemonService

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/Tonipenyallop/pokedex-api/constants"
	pokemonServiceHelper "github.com/Tonipenyallop/pokedex-api/helpers"
	pokemonRepository "github.com/Tonipenyallop/pokedex-api/repository"
	"github.com/Tonipenyallop/pokedex-api/types"
	"github.com/patrickmn/go-cache"
	"google.golang.org/api/youtube/v3"
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
		return *cachedGenPokemons, nil
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
		return nil, fmt.Errorf("failed to fetch pokemons by generation: %v", err)
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
		return nil, fmt.Errorf("failed to get flavor text:%v", err)
	}

	tmp, err := pokemonRepository.GetEvolutionChain(flavorText.EvolutionChain.URL)

	// taesu here
	// // create helper to get array of 3 pokemons
	if err != nil {
		return nil, fmt.Errorf("failed to fetch evolution chain detail:%v", err)
	}
	fmt.Println("tmp", tmp)

	// evolutionChain, err := pokemonRepository.GetEvolutionChainPokemonNames(pokemonId)
	// if err != nil {
	// 	return nil, fmt.Errorf("Failed to get evolution chain")
	// }

	// [types.SpeciesInfo, ]

	return flavorText, nil

}

func GetPokemonDetail(pokemonId string) (*pokemonRepository.TmpPokemon, error) {
	convertedPokemonId, err := strconv.Atoi(pokemonId)
	if err != nil {
		return nil, fmt.Errorf("failed to convert pokemonId:%v", err)
	}

	// check gen cache first
	genId := pokemonServiceHelper.GetGenIdByPokemonId(convertedPokemonId)

	// get pokemon from gen cache
	genPokemonsInCache := pokemonServiceHelper.GetPokemonsFromCacheByGen(genId, pokemonCache)

	if genPokemonsInCache != nil {
		for _, pokemonInCache := range *genPokemonsInCache {
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
		return nil, fmt.Errorf("failed to get pokemon detail %v", err)
	}

	return pokemon, nil

}

func GetYoutubePlayList(youtubeService *youtube.Service) (*youtube.PlaylistItemListResponse, error) {
	partList := []string{"snippet,contentDetails"}

	playlistId := os.Getenv("PLAYLIST_ID")
	// assumes the playlist is 15 length
	call := youtubeService.PlaylistItems.List(partList).PlaylistId(playlistId).MaxResults(constants.MAX_PLAYLIST_COUNT)
	res, err := call.Do()
	if err != nil {
		return nil, fmt.Errorf("failed to search:%v", err)
	}
	return res, nil

}

func GetYoutubeDescriptionById(youtubeService *youtube.Service, videoId string) (*types.GetYoutubeDescriptionByIdResponse, error) {
	call := youtubeService.Videos.List([]string{"snippet"}).Id(videoId)

	res, err := call.Do()

	if err != nil {
		return nil, fmt.Errorf("failed to fetch video by videoId:%v", err)
	}

	musicDescription := res.Items[0].Snippet.Localized.Description
	pokemonMusicDescriptions := pokemonServiceHelper.HelperDescription(musicDescription)
	musicId := res.Items[0].Id

	output := types.GetYoutubeDescriptionByIdResponse{
		MusicDescription: pokemonMusicDescriptions,
		MusicId:          musicId,
	}

	return &output, nil
}
