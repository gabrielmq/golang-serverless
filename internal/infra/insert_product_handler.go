package infra

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gabrielmq/golang-serverless/internal/entity"
	"github.com/gabrielmq/golang-serverless/internal/infra/dynamo"
	"github.com/google/uuid"
)

func InsertProductHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var product entity.Product
	if err := json.Unmarshal([]byte(request.Body), &product); err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	// inicializa o servico do dynamo
	svc := dynamo.NewDynamoSession()

	// preparando o dado para insert no dynamo
	product.ID = uuid.New().String()
	input := &dynamodb.PutItemInput{
		TableName: aws.String("Products"),
		Item: map[string]*dynamodb.AttributeValue{
			"id":    {S: aws.String(product.ID)},
			"name":  {S: aws.String(product.Name)},
			"price": {N: aws.String(fmt.Sprintf("%f", product.Price))},
		},
	}

	// inserindo na tabela do dynamo
	_, err := svc.PutItem(input)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	body, err := json.Marshal(product)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}, nil
}
