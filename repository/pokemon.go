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
}


func GetAllPokemons ()([]TmpPokemon, error){

	svc, err := pokemonDynamo.GetDynamo()
	if err != nil {
		return nil, fmt.Errorf("Failed to get dynamo:",err)
	}

	// Use a dynamically growing slice for storing all Pokémons
	var pokemons []TmpPokemon

	input := &dynamodb.ScanInput{
		TableName: aws.String("Pokemons"),
	}

	for {
		// Perform the Scan operation
		result, err := svc.Scan(input)
		if err != nil {
			return nil, fmt.Errorf("Failed to scan table: %v", err)
		}

		// Process the items in the current batch
		for _, item := range result.Items {

			id, err := strconv.Atoi(*item["ID"].N)
			if err != nil {
				log.Printf("Failed to convert ID: %v", err)
				continue
			}

			name := *item["Name"].S


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

			pokemons = append(pokemons, TmpPokemon{
				ID:      id,
				Name:    name,
				Sprites: sprites,
				Types:   types,
			})

		}

		// Check if there are more items to fetch
		if result.LastEvaluatedKey == nil {
			break // Exit the loop if no more items
		}

		// Update the Scan input with the LastEvaluatedKey for the next batch
		input.ExclusiveStartKey = result.LastEvaluatedKey
	}
	return pokemons, nil
}

