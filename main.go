package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	pokemonService "github.com/Tonipenyallop/pokedex-api/services"

	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// CORS setting
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin", "Content-type", "Authorization"},
	}))

	// Define routes
	pokemon := router.Group("/pokemon")
	{
		pokemon.GET("/all", getAllPokemons)
		pokemon.GET("/gen/:generationId", getPokemonsByGen)
		pokemon.GET("/evolution-chain/:pokemonId", getPokemonFlavorTextAndEvolutionChain)
		pokemon.GET("/music/:musicIndex", getMusicDescriptions)
		pokemon.GET("/:pokemonId", getPokemonDetails) // wildcard last to avoid conflicts
	}

	// Start the server
	router.Run(":8080")
}

func getMusicDescriptions(c *gin.Context) {
	godotenv.Load() // optional: env vars may come from container env
	musicIndex, found := c.Params.Get("musicIndex")
	if !found {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "failed to get musicIndex from parameter"})
	}

	convertedIndex, err := strconv.Atoi(musicIndex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("failed to convert musicIndex: %v", musicIndex)})
	}

	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(os.Getenv("YOUTUBE_API_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("failed to initiate new youtube service: %v", err)})
	}

	playlist, err := pokemonService.GetYoutubePlayList(youtubeService)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to fetch youtube playlist: %s", err)})
	}

	videoId := playlist.Items[convertedIndex].Snippet.ResourceId.VideoId

	description, err := pokemonService.GetYoutubeDescriptionById(youtubeService, videoId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("failed to get youtube description by id: %v", err)})
	}
	c.IndentedJSON(http.StatusOK, description)
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
