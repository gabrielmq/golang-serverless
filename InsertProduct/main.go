package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func InsertProduct(ctx context.Context, request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	var product Product
	if err := json.Unmarshal([]byte(request.Body), &product); err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}
	}

	// abre uma sessao com a aws
	sess := session.Must(session.NewSession())

	// inicializa o servico do dynamo
	svc := dynamodb.New(sess)

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
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}
	}

	body, err := json.Marshal(product)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}
}

func main() {
	// iniciando a lambda function
	lambda.Start(InsertProduct)
}
