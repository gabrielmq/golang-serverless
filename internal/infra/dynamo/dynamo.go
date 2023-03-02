package dynamo

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewDynamoSession() *dynamodb.DynamoDB {
	cfg := &aws.Config{
		Region:     aws.String(os.Getenv("AWS_DEFAULT_REGION")),
		Endpoint:   aws.String(os.Getenv("LOCALSTACK_HOSTNAME") + ":4566"),
		DisableSSL: aws.Bool(true),
	}

	// abre uma sessao com a aws
	sess := session.Must(session.NewSession(cfg))

	// inicializa o servico do dynamo
	return dynamodb.New(sess)
}
