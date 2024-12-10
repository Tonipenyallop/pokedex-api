package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	pokemonServiceHelper "github.com/Tonipenyallop/pokedex-api/helpers"
	pokemonService "github.com/Tonipenyallop/pokedex-api/services"
	"github.com/Tonipenyallop/pokedex-api/types"

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
	router.GET("/pokemon/music", tmpMusic)

	// Start the server
	router.Run("localhost:8080")
}


func tmpMusic(c * gin.Context){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env vars",err)
	}
	ctx := 	context.Background() 
	yts, err := youtube.NewService(ctx,option.WithAPIKey(os.Getenv("YOUTUBE_API_KEY")))

	tmp := []string {"snippet,contentDetails,statistics"}

	videoID := "DzFlPk2UnQg"

    // Request video details
    call := yts.Videos.List(tmp).Id(videoID)

	res, err := call.Do()
	if err != nil {
		log.Fatal("Failed to search",err)
	}


	musicDescription := res.Items[0].Snippet.Localized.Description
	pokemonMusicDescriptions :=pokemonServiceHelper.HelperDescription(musicDescription)
	musicId := res.Items[0].Id
	type Musica struct {
		MusicDescription []types.YoutubeMusic `json:"musicDescription"`
		MusicId string `json:"musicId"`
	}
	 tmpMusica :=  Musica{
		MusicDescription: pokemonMusicDescriptions,
		MusicId: musicId,
	}

	c.IndentedJSON(http.StatusOK, tmpMusica)
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
