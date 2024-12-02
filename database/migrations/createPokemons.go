package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/Tonipenyallop/pokedex-api/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/joho/godotenv"
)


func determineGeneration(id int) string {
    switch {
    case id >= 1 && id <= 151:
        return "first"
    case id >= 152 && id <= 251:
        return "second"
    case id >= 252 && id <= 386:
        return "third"
    case id >= 387 && id <= 493:
        return "fourth"
    case id >= 494 && id <= 649:
        return "fifth"
    case id >= 650 && id <= 721:
        return "sixth"
    case id >= 722 && id <= 809:
        return "seventh"
    case id >= 810 && id <= 905:
        return "eighth"
    case id >= 906 && id <= 1025:
        return "ninth"
    default:
        return "Unknown Generation"
    }
}

func serializeToJSON(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to serialize data to JSON: %v", err)
	}
	return string(jsonData)
}

func fetchPokemonData(id int) (map[string]*dynamodb.AttributeValue, error) {
	formattedURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", id)
	resp, err := http.Get(formattedURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Pokémon data for ID %d: %w", id, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body for ID %d: %w", id, err)
	}

	var pokemon types.Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal Pokémon data for ID %d: %w", id, err)
	}
	gen := determineGeneration(pokemon.ID)

	item := map[string]*dynamodb.AttributeValue{
		"ID":              {N: aws.String(fmt.Sprintf("%d", pokemon.ID))},
		"Name":            {S: aws.String(pokemon.Name)},
		"Abilities":       {S: aws.String(serializeToJSON(pokemon.Abilities))},
		"Cries":           {S: aws.String(serializeToJSON(pokemon.Cries))},
		"Forms":           {S: aws.String(serializeToJSON(pokemon.Forms))},
		"Height":          {N: aws.String(fmt.Sprintf("%d", pokemon.Height))},
		"Species":         {S: aws.String(serializeToJSON(pokemon.Species))},
		"Sprites":         {S: aws.String(serializeToJSON(pokemon.Sprites))},
		"Types":           {S: aws.String(serializeToJSON(pokemon.Types))},
		"Weight":          {N: aws.String(fmt.Sprintf("%d", pokemon.Weight))},
		"Generation":      {S: aws.String(gen)},
	}

	return item, nil
}

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env vars")
		return
	}

	awsProfile := os.Getenv("AWS_PROFILE_NAME")
	awsRegion := os.Getenv("AWS_REGION")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: awsProfile,
		Config:  aws.Config{Region: aws.String(awsRegion)},
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	fromPokemonIndex := 1
	lastPokemonIndex := 1025
	batchSize := 10 // Number of Pokémon to batch

	var wg sync.WaitGroup
	mutex := sync.Mutex{}

	itemsBatch := make([]map[string]*dynamodb.AttributeValue, 0, batchSize)

	// Fetch and batch Pokémon data
	for i := fromPokemonIndex; i <= lastPokemonIndex; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			item, err := fetchPokemonData(id)
			if err != nil {
				log.Printf("Error fetching data for ID %d: %v", id, err)
				return
			}

			mutex.Lock()
			itemsBatch = append(itemsBatch, item)
			if len(itemsBatch) >= batchSize {
				writeBatch(svc, itemsBatch)
				itemsBatch = itemsBatch[:0] // Clear the batch
			}
			mutex.Unlock()
		}(i)
	}

	wg.Wait()

	// Write any remaining items
	if len(itemsBatch) > 0 {
		writeBatch(svc, itemsBatch)
	}

	log.Println("Successfully added all items.")
}

func writeBatch(svc *dynamodb.DynamoDB, items []map[string]*dynamodb.AttributeValue) {
	for {
		requestItems := make(map[string][]*dynamodb.WriteRequest)

		for _, item := range items {
			requestItems["Pokemons"] = append(requestItems["Pokemons"], &dynamodb.WriteRequest{
				PutRequest: &dynamodb.PutRequest{
					Item: item,
				},
			})
		}

		input := &dynamodb.BatchWriteItemInput{
			RequestItems: requestItems,
		}

		_, err := svc.BatchWriteItem(input)
		if err != nil {
			log.Printf("Error writing batch: %v. Retrying...", err)
			time.Sleep(1 * time.Second)
			continue
		}

		break
	}
}