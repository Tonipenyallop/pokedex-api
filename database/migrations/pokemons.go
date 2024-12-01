package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/joho/godotenv"
)



func main(){

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env vars")
		return
	}


	awsProfile := os.Getenv("AWS_PROFILE_NAME")
	awsRegion := os.Getenv("AWS_REGION")


	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:  awsProfile,
		Config: aws.Config{Region: aws.String(awsRegion)},
	}))


	// Create dynamoDB client
	svc := dynamodb.New(sess)



	item := map[string]*dynamodb.AttributeValue{
		"ID": {
			N: aws.String("1"),
		},
		"Name": {
			S: aws.String("Bulbasaur"),
		},
		"Type": {
			S: aws.String("Grass/Poison"),
		},
		"BaseExperience": {
			N: aws.String("64"),
		},
		"Height": {
			N: aws.String("7"),
		},
		"Weight": {
			N: aws.String("69"),
		},
	}


	// create the PutItemInput for the request
	input := &dynamodb.PutItemInput{
		TableName: aws.String("Pokemons"),
		Item: item,
	}


	// add the item to the table
	_, err = svc.PutItem(input)

	if err != nil {
		log.Fatal("Failed to add item", input, err)
		return
	}


	fmt.Printf("Successfully add all items")




}



