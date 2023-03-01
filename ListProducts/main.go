package main

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func ListProducts(ctx context.Context, request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	// preparando a busca no dynamo
	input := &dynamodb.ScanInput{
		TableName: aws.String("Products"),
	}

	result, err := svc.Scan(input)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}
	}

	var products []Product
	for _, item := range result.Items {
		price, err := strconv.Atoi(*item["price"].N)
		if err != nil {
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}
		}

		products = append(products, Product{
			ID:    *item["id"].S,
			Name:  *item["name"].S,
			Price: float64(price),
		})
	}

	body, err := json.Marshal(products)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}
}

func main() {
	lambda.Start(ListProducts)
}