package infra

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gabrielmq/golang-serverless/internal/entity"
	"github.com/gabrielmq/golang-serverless/internal/infra/dynamo"
)

func ListProductsHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	svc := dynamo.NewDynamoSession()

	// preparando a busca no dynamo
	input := &dynamodb.ScanInput{
		TableName: aws.String("Products"),
	}

	result, err := svc.Scan(input)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	var products []entity.Product
	for _, item := range result.Items {
		price, err := strconv.Atoi(*item["price"].N)
		if err != nil {
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		}

		products = append(products, entity.Product{
			ID:    *item["id"].S,
			Name:  *item["name"].S,
			Price: float64(price),
		})
	}

	body, err := json.Marshal(products)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}, nil
}
