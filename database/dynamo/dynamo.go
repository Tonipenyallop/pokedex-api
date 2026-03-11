package pokemonDynamo

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/joho/godotenv"
)

func GetDynamo() (*dynamodb.DynamoDB, error) {

	godotenv.Load() // optional: env vars may come from container env

	
	cfg := &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}

	if endpoint := os.Getenv("DYNAMO_ENDPOINT"); endpoint != "" {
		cfg.Endpoint = aws.String(endpoint)
	}

	sess, err := session.NewSession(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	svc := dynamodb.New(sess)

	return svc, nil

}
