package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"

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

// func getPokemonDetails(c *gin.Context) {
// 	fmt.Println("getPokemonDetail was called")
// 	pokemonId, booleanValue := c.Params.Get("pokemonId")

// 	if !booleanValue {
// 		log.Fatal("Failed to get pokemonId param")
// 		return
// 	}

// 	isRange := c.Query("isRange")

// 	// want single data. Should rename it!
// 	if isRange == "true" {

// 		formattedUrl := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonId)

// 		resp, err := http.Get(formattedUrl)
// 		if err != nil {
// 			log.Fatal("Failed to fetch pokemon ditto")
// 			return
// 		}
// 		defer resp.Body.Close()

// 		// maybe not use ioutil.ReadAll?
// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			log.Fatalf("Failed to read response body: %v", err)
// 			return
// 		}

// 		// Decode JSON
// 		var pokemon types.Pokemon
// 		err = json.Unmarshal(body, &pokemon)
// 		if err != nil {
// 			log.Fatalf("Error decoding JSON: %v", err)
// 			return
// 		}

		
// 		// retuning single pokemon as an array to match same data type
// 		tmp := [1]types.Pokemon{}
// 		// tmp[0] = types.Pokemon(pokemonDetail)
// 		tmp[0] = types.Pokemon(pokemon)
// 		// fmt.Println("tmp[0]",tmp[0])
// 		c.IndentedJSON(http.StatusOK, tmp)
// 		return

// 	}

// 	convertedId, err := strconv.Atoi(pokemonId)
// 	if err != nil {
// 		log.Fatal("Failed to convert string to int")
// 	}

// 	var wg sync.WaitGroup
// 	pokemons := make([]types.Pokemon, convertedId)
// 	// pokemons := make([]types.Pokemon, convertedId)

// 	for i := 0; i < convertedId; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			formattedUrl := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", i + 1)
			
// 			pokemonDetail, err2 := http.Get(formattedUrl)
// 			// pokemonDetail, err2 := pokeapi.Pokemon(strconv.Itoa(i + 1))
// 			if err2 != nil {
// 				log.Fatal("Failed to fetch pokemonDetail", err2)
// 			}

// 			body, err := io.ReadAll(pokemonDetail.Body)
// 		if err != nil {
// 			log.Fatalf("Failed to read response body: %v", err)
// 			return
// 		}

// 		// Decode JSON
// 		var pokemon types.Pokemon
// 		err = json.Unmarshal(body, &pokemon)
// 		if err != nil {
// 			log.Fatalf("Error decoding JSON: %v", err)
// 			return
// 		}

// 			pokemons[i] = types.Pokemon(pokemon)
// 		}()

// 	}

// 	// Need to wait
// 	wg.Wait()

// 	c.IndentedJSON(http.StatusOK, pokemons)

// }

func getPokemonDetailsByGeneration(c *gin.Context) {
	fmt.Println("getPokemonDetailsByGeneration was called")
	generationId, booleanValue := c.Params.Get("generationId")
	if !booleanValue {
		log.Fatal("Failed to get generationId from Param")
	}
	from := constants.GENERATION_MAP[generationId][0]
	to := constants.GENERATION_MAP[generationId][1]
	fmt.Println("from", from, "to", to)

	var wg sync.WaitGroup

	var pokemons = make([]types.Pokemon, to-from+1)

	// 1 index
	for i := from; i <= to; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pokemon, err := pokeapi.Pokemon(strconv.Itoa(i))

			if err != nil {
				log.Fatal("Failed to fetch pokemon", err)
			}

			pokemons[i-from] = types.Pokemon(pokemon)

		}()
	}

	wg.Wait()

	c.IndentedJSON(http.StatusOK, pokemons)

}
