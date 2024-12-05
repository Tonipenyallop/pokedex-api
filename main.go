package main

import (
	"fmt"
	"net/http"

	pokemonService "github.com/Tonipenyallop/pokedex-api/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	router.GET("/pokemon/all", getAllPokemons)
	router.GET("/pokemon/:pokemonId", getPokemonDetails)
	router.GET("/pokemon/gen/:generationId", getPokemonsByGen)
	// router.GET("/pokemon/evolution-chain/:pokemonId", getEvolutionChainById)
	router.GET("/pokemon/evolution-chain/:pokemonId", getPokemonFlavorTextAndEvolutionChain)

	// Start the server
	router.Run("localhost:8080")
}

func getAllPokemons(c *gin.Context) {
	pokemons, err := pokemonService.GetAllPokemons()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, pokemons)
}

func getPokemonDetails(c *gin.Context) {
	// Get pokemonId from URL parameter
	pokemonId, found := c.Params.Get("pokemonId")
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pokemonId parameter is required"})
		return
	}

	pokemon, err := pokemonService.GetPokemonDetail(pokemonId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": fmt.Sprintf("Failed to fetch pokemon detail %s", err)})
	}

	c.IndentedJSON(http.StatusOK, pokemon)

}

func getPokemonsByGen(c *gin.Context) {

	// // Get generationId from URL parameter
	generationId, found := c.Params.Get("generationId")
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "generationId parameter is required"})
		return
	}

	pokemons, err := pokemonService.GetPokemonsByGen(generationId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, pokemons)
}

// this method will return up to 3 pokemons of array
// func getEvolutionChainById(c *gin.Context){
// 	pokemonId, found := c.Params.Get("pokemonId")

// 	if !found {
// 		c.JSON(http.StatusBadRequest, gin.H{"error" : "pokemonId param is required"})
// 	}

// 	evolutionChain, err := pokemonService.GetEvolutionChainById(pokemonId)

// 	if err != nil {
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error" : fmt.Sprint("failed to get evolution chain with id:%s %s",pokemonId,err)})
// 	}

// 	c.IndentedJSON(http.StatusOK, evolutionChain)

// }

func getPokemonFlavorTextAndEvolutionChain(c *gin.Context) {
	pokemonId, found := c.Params.Get("pokemonId")

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pokemonId param is required"})
	}

	evolutionChain, err := pokemonService.GetPokemonFlavorTextAndEvolutionChain(pokemonId)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint("failed to get evolution chain with id:", err)})
	}

	c.IndentedJSON(http.StatusOK, evolutionChain)

}
