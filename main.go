package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	// Start the server
	router.Run("localhost:8080")
}

// Function to fetch Pokémon data and return JSON
func getPokemons(c *gin.Context) {
	fmt.Println("BE getPokemons are called mate")
	// URL for the PokéAPI
	url := "https://pokeapi.co/api/v2/pokemon/ditto"

	// Perform HTTP GET request
	resp, err := http.Get(url)


	if err != nil {
		fmt.Println("Error occurred while fetching data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from PokéAPI"})
		return
	}
	defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error occurred while reading response:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read API response"})
		return
	}

	// Parse the JSON response to validate it
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error occurred while parsing JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON"})
		return
	}

	// Return the entire response as JSON
	c.IndentedJSON(http.StatusOK, data)
}