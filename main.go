package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"

	pokemonService "github.com/Tonipenyallop/pokedex-api/services"

	"github.com/Tonipenyallop/pokedex-api/constants"
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
	router.GET("/pokemon/all", getAllPokemons)
	router.GET("/pokemon/:pokemonId", getPokemonDetails)
	router.GET("/pokemon/gen/:generationId", getPokemonDetailsByGeneration)

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




func getAllPokemons(c *gin.Context) {
	pokemons, err := pokemonService.GetAllPokemons()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,err)
		return 
	}

	c.IndentedJSON(http.StatusOK, pokemons)
}

func getPokemonDetails(c *gin.Context) {
	fmt.Println("getPokemonDetail was called")

	// Get pokemonId from URL parameter
	pokemonId, found := c.Params.Get("pokemonId")
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pokemonId parameter is required"})
		return
	}

	isRange := c.Query("isRange")

	// If single data request
	if isRange == "true" {
		formattedUrl := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonId)

		resp, err := http.Get(formattedUrl)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Pokemon details"})
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
			return
		}

		var pokemon types.Pokemon
		err = json.Unmarshal(body, &pokemon)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode JSON"})
			return
		}

		// Return single Pokemon as an array for consistency
		c.IndentedJSON(http.StatusOK, []types.Pokemon{pokemon})
		return
	}

	// Handle range data
	convertedId, err := strconv.Atoi(pokemonId)
	if err != nil || convertedId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pokemonId"})
		return
	}

	var wg sync.WaitGroup

	// Pre-allocate a slice for ordered insertion
	pokemons := make([]types.Pokemon, convertedId)

	// Rate-limiting channel
	rateLimiter := make(chan struct{}, 10) // Allow up to 10 concurrent requests

	for i := 1; i <= convertedId; i++ {
		wg.Add(1)
		rateLimiter <- struct{}{} // Acquire a slot

		go func(index int) {
			defer wg.Done()
			defer func() { <-rateLimiter }() // Release the slot

			formattedUrl := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", index)
			resp, err := http.Get(formattedUrl)
			if err != nil {
				fmt.Printf("Failed to fetch Pokemon details for ID %d: %v\n", index, err)
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Failed to read response body for ID %d: %v\n", index, err)
				return
			}

			var pokemon types.Pokemon
			err = json.Unmarshal(body, &pokemon)
			if err != nil {
				fmt.Printf("Failed to decode JSON for ID %d: %v\n", index, err)
				return
			}

			// Insert Pokemon into the correct index
			pokemons[index-1] = pokemon
		}(i)
	}

	wg.Wait()

	c.IndentedJSON(http.StatusOK, pokemons)
}

func getPokemonDetailsByGeneration(c *gin.Context) {
	fmt.Println("getPokemonDetailsByGeneration was called")

	// Get generationId from URL parameter
	generationId, found := c.Params.Get("generationId")
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "generationId parameter is required"})
		return
	}

	// Fetch the range for the given generation
	rangeValues, exists := constants.GENERATION_MAP[generationId]
	if !exists || len(rangeValues) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid generationId"})
		return
	}

	from, to := rangeValues[0], rangeValues[1]
	fmt.Println("from", from, "to", to)

	// Pre-allocate slice to hold Pokémon data in order
	pokemons := make([]types.Pokemon, to-from+1)

	var wg sync.WaitGroup

	// Rate-limiting channel for controlled concurrency
	rateLimiter := make(chan struct{}, 10) // Limit to 10 concurrent requests

	for i := from; i <= to; i++ {
		wg.Add(1)
		rateLimiter <- struct{}{} // Acquire a slot

		go func(index int) {
			defer wg.Done()
			defer func() { <-rateLimiter }() // Release the slot

			// Fetch Pokémon data using the given endpoint
			url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", index)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Failed to fetch Pokémon for ID %d: %v\n", index, err)
				return
			}
			defer resp.Body.Close()

			// Read response body
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Failed to read response body for ID %d: %v\n", index, err)
				return
			}

			// Decode JSON
			var pokemon types.Pokemon
			err = json.Unmarshal(body, &pokemon)
			if err != nil {
				fmt.Printf("Failed to decode JSON for ID %d: %v\n", index, err)
				return
			}

			// Store in the correct index
			pokemons[index-from] = pokemon
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	c.IndentedJSON(http.StatusOK, pokemons)
}
