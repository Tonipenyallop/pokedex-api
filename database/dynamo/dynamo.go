package pokemonDynamo

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/joho/godotenv"
)


func GetDynamo()(*dynamodb.DynamoDB,error) {

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Failed to load env vars:",err)
	}
	
	awsProfile := os.Getenv("AWS_PROFILE_NAME")
	awsRegion := os.Getenv("AWS_REGION")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: awsProfile,
		Config: aws.Config{
			Region: aws.String(awsRegion),
		},
	}))

	svc := dynamodb.New(sess)

	return svc , nil


}