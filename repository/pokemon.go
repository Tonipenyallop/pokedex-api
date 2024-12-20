package pokemonRepository

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	pokemonDynamo "github.com/Tonipenyallop/pokedex-api/database/dynamo"
	"github.com/Tonipenyallop/pokedex-api/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// somehow cannot use GetAllPokemonsResponse
type TmpPokemon struct {
	ID      int              `json:"id"`
	Name    string           `json:"name"`
	Sprites types.Sprites    `json:"sprites"`
	Types   []types.TypeSlot `json:"types"`
	Generation   string `json:"generation"`
	Cries struct{
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
}

func GetAllPokemons() ([]TmpPokemon, error) {

	svc, err := pokemonDynamo.GetDynamo()
	if err != nil {
		return nil, fmt.Errorf("Failed to get dynamo:", err)
	}

	// Use a dynamically growing slice for storing all Pokémons
	var pokemons []TmpPokemon

	input := &dynamodb.ScanInput{
		TableName: aws.String("Pokemons"),
	}

	for {
		// Perform the Scan operation
		res, err := svc.Scan(input)
		if err != nil {
			return nil, fmt.Errorf("Failed to scan table: %v", err)
		}

		pokemons = convertPokemonPropsFromAWS(res.Items,pokemons)

		// Check if there are more items to fetch
		if res.LastEvaluatedKey == nil {
			break // Exit the loop if no more items
		}

		// Update the Scan input with the LastEvaluatedKey for the next batch
		input.ExclusiveStartKey = res.LastEvaluatedKey
	}
	return pokemons, nil
}



func GetPokemonsByGen(genId string) ([]TmpPokemon, error) {


	// Initialize DynamoDB client
	svc, err := pokemonDynamo.GetDynamo()
	if err != nil {
		return nil, fmt.Errorf("Failed to get DynamoDB client: %w", err)
	}


	// Define Scan input
	input := &dynamodb.ScanInput{
		TableName:        aws.String("Pokemons"),
		FilterExpression: aws.String("Generation = :gen"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":gen": {S: aws.String(genId)},
		},
	}


	var pokemons []TmpPokemon
	// Execute Scan

	for {
		res, err := svc.Scan(input)
		if err != nil {
			return nil, fmt.Errorf("Failed to scan DynamoDB: %w", err)
		}

		pokemons = convertPokemonPropsFromAWS(res.Items,pokemons)
		// Check if there are more items to fetch
		if res.LastEvaluatedKey == nil {
			break // Exit the loop if no more items
		}

		// Update the Scan input with the LastEvaluatedKey for the next batch
		input.ExclusiveStartKey = res.LastEvaluatedKey
	}

	return pokemons, nil
}

func convertPokemonPropsFromAWS(items []map[string]*dynamodb.AttributeValue, array []TmpPokemon) []TmpPokemon{
	for _, item := range items {

		id, err := strconv.Atoi(*item["ID"].N)
		if err != nil {
			log.Printf("Failed to convert ID: %v", err)
			continue
		}

		name := *item["Name"].S

		generation:= *item["Generation"].S

		var sprites types.Sprites
		err = json.Unmarshal([]byte(*item["Sprites"].S), &sprites)
		if err != nil {
			log.Printf("Failed to unmarshal Sprites for ID %d: %v", id, err)
			continue
		}

		
		// types collision to native name
		var typeVar []types.TypeSlot
		err = json.Unmarshal([]byte(*item["Types"].S), &typeVar)
		if err != nil {
			log.Printf("Failed to unmarshal Types for ID %d: %v", id, err)
			continue
		}


		var cries types.Cries
		err = json.Unmarshal([]byte(*item["Cries"].S), &cries)
		if err != nil {
			log.Printf("Failed to unmarshal Cries for ID %d: %v", id, err)
			continue
		}


		array = append(array, TmpPokemon{
			ID:      id,
			Name:    name,
			Sprites: sprites,
			Types:   typeVar,
			Generation: generation,
			Cries: cries,
		})
	}

	return array


}


type RevolutionChain struct {

}

// get from pokemon species
// seems like an ID of endpoint is not corresponding to pokemon ID
func GetEvolutionChain(url string)(*[]int, error){
	// url := fmt.Sprintf("https://pokeapi.co/api/v2/evolution-chain/%s", pokemonId)
	fmt.Println("url hehe",url)
	res, err :=  http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch evolution chain",err)
	}

	defer res.Body.Close()


	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to read all body",err)
	}


	var evolutionChain types.EvolutionChain

	err = json.Unmarshal(body,&evolutionChain)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body",err)
	}

	

	var result []int
	getIDsFromEvolutionChain(evolutionChain.Chain, &result)


	return &result , nil

}


func GetPokemonFlavorTextAndEvolutionChain(pokemonId string)(*types.SpeciesInfo,error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-species/%s", pokemonId)

	res, err :=  http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch evolution chain",err)
	}

	defer res.Body.Close()	


	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to read all body",err)
	}


	var speciesInfo types.SpeciesInfo

	err = json.Unmarshal(body,&speciesInfo)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body",err)
	}

	return &speciesInfo, nil

}


func GetPokemonDetail(pokemonId string)(*TmpPokemon, error){
	fmt.Println("repository was called amte")
	formattedUrl := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonId)

	resp, err := http.Get(formattedUrl)
	if err != nil { 
		return nil, fmt.Errorf("Failed to fetch Pokemon details",err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil { 
		return nil, fmt.Errorf("Failed to read response body",err)
	}

	var pokemon TmpPokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil { 
		return nil, fmt.Errorf("Failed to decode JSON",err)
	}	


	return &pokemon, nil
}

// helper

func extractID(url string) int {
	parts := strings.Split(strings.TrimSuffix(url, "/"), "/")
	id := parts[len(parts)-1]
	parsedID := 0
	fmt.Sscanf(id, "%d", &parsedID)
	return parsedID
}


func getIDsFromEvolutionChain(chain types.Chain, result *[]int) {
	*result = append(*result, 
		  extractID(chain.Species.URL),
	)

	for _, nextChain := range chain.EvolvesTo {
		getIDsFromEvolutionChain(nextChain, result)
	}
}