package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mtslzr/pokeapi-go"
)

// Struct to represent Pokémon details from PokéAPI
type PokemonAPIResponse struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
	Height int    `json:"height"`
}

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Cors setting
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET","POST"},
		AllowHeaders: []string{"Origin","Content-type","Authorization"},
	}))
	

	// Define routes
	router.GET("/pokemon", getPokemons)
	router.GET("/pokemon/:pokemonId", getPokemonDetail)


	// Start the server
	router.Run("localhost:8080")
}

// Function to fetch Pokémon data and return JSON
func getPokemons(c *gin.Context) {
	fmt.Println("BE getPokemons are called mate")
	// URL for the PokéAPI
	pokemons, err := pokeapi.Resource("pokemon",0,151)

	if (err != nil){
		panic("failed to fetch pokemons")
	}


	c.IndentedJSON(http.StatusOK,pokemons)
}


func getPokemonDetail(c *gin.Context){
	fmt.Println("getPokemonDetail was called")
	pokemonId, booleanValue := c.Params.Get("pokemonId")
	
	if (!booleanValue){
		log.Fatal("Failed to get pokemonId param")
	}

	 pokemonDetail, err2 := pokeapi.Pokemon(pokemonId)

	 if (err2 != nil){
		log.Fatal("Failed to fetch pokemonDetail",err2)
	 }


	 c.IndentedJSON(http.StatusOK, pokemonDetail)
	 
	
}



