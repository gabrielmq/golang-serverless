package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gabrielmq/golang-serverless/internal/infra"
)

func main() {
	lambda.Start(infra.ListProductsHandler)
}
