package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/Tonipenyallop/pokedex-api/types"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mtslzr/pokeapi-go"
)
func main() {
	// Initialize Gin router
	router := gin.Default()

	// Cors setting
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin", "Content-type", "Authorization"},
	}))

	// Define routes
	router.GET("/pokemon", getPokemons)
	router.GET("/pokemon/:pokemonId", getPokemonDetails)

	// Start the server
	router.Run("localhost:8080")
}

// Function to fetch Pokémon data and return JSON
func getPokemons(c *gin.Context) {
	fmt.Println("BE getPokemons are called mate")
	// URL for the PokéAPI
	pokemons, err := pokeapi.Resource("pokemon", 0, 151)

	if err != nil {
		panic("failed to fetch pokemons")
	}

	c.IndentedJSON(http.StatusOK, pokemons)
}

func getPokemonDetails(c *gin.Context) {
	fmt.Println("getPokemonDetail was called")
	pokemonId, booleanValue := c.Params.Get("pokemonId")

	if !booleanValue {
		log.Fatal("Failed to get pokemonId param")
	}

	convertedId, err := strconv.Atoi(pokemonId)
	if err != nil {
		log.Fatal("Failed to convert string to int")
	}

	var wg sync.WaitGroup
	pokemons := make([]types.Pokemon, convertedId)

	for i := 0; i < convertedId; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pokemonDetail, err2 := pokeapi.Pokemon(strconv.Itoa(i + 1))
			if err2 != nil {
				log.Fatal("Failed to fetch pokemonDetail", err2)
			}
			pokemons[i] = types.Pokemon(pokemonDetail)
		}()

	}

	// Need to wait
	wg.Wait()

	c.IndentedJSON(http.StatusOK, pokemons)

}
