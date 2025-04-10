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

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Failed to load env vars:", err)
	}

	awsRegion := os.Getenv("AWS_REGION")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	svc := dynamodb.New(sess)

	return svc, nil

}
