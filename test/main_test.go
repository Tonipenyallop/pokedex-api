package tests

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	pokemonServiceHelper "github.com/Tonipenyallop/pokedex-api/helpers"
	pokemonRepository "github.com/Tonipenyallop/pokedex-api/repository"
	"github.com/patrickmn/go-cache"
)

func TestGetPokemonsFromCacheByGen(t *testing.T) {

	var pokemonCache = cache.New(24*time.Hour, 24*time.Hour)

	t.Cleanup(func() {
		pokemonCache.Flush()
	})

	genId := "first"

	cacheKey := pokemonServiceHelper.GetGenCacheKey(genId)

	type FakePokemon struct {
		ID   string
		Name string
	}

	pokemons := []pokemonRepository.TmpPokemon{
		{
			ID:   9,
			Name: "Toni",
		},
	}

	pokemonCache.Set(cacheKey, pokemons, 24*time.Hour)

	res := pokemonServiceHelper.GetPokemonsFromCacheByGen(genId, pokemonCache)
	fmt.Println("reflect.DeepEqual(res,pokemons)", reflect.DeepEqual(res, pokemons))
	if !reflect.DeepEqual(res, pokemons) {
		t.Fatalf("expect res: and pokemons: to be same")
	}

}




func TestGetGenIdByPokemonId(t *testing.T){
    testCases := []struct{
        pokemonId int
        expect string
    }{
        {151,"first"},
        {152,"second"},
        {300,"third"},
        {400,"fourth"},
        {500,"fifth"},
        {700,"sixth"},
        {750,"seventh"},
        {811,"eighth"},
        {907,"ninth"},
    }

    for _, testCase := range testCases{
        res := pokemonServiceHelper.GetGenIdByPokemonId(testCase.pokemonId)

        if res != testCase.expect{
            t.Fatalf("expect %s to be %s",res,testCase.expect)
        }
        
    }

}
