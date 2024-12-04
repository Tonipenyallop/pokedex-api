package pokemonRepository

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

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

		var types []types.TypeSlot
		err = json.Unmarshal([]byte(*item["Types"].S), &types)
		if err != nil {
			log.Printf("Failed to unmarshal Types for ID %d: %v", id, err)
			continue
		}

		array = append(array, TmpPokemon{
			ID:      id,
			Name:    name,
			Sprites: sprites,
			Types:   types,
			Generation: generation,
		})
	}

	return array


}