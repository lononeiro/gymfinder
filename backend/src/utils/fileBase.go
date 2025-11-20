package utils

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Cria e retorna um client S3 configurado para Filebase
func NewFilebaseClient(ctx context.Context) (*s3.Client, error) {

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			os.Getenv("FILEBASE_S3_ACCESS_KEY"),
			os.Getenv("FILEBASE_S3_SECRET_KEY"),
			"",
		)),
		config.WithRegion(os.Getenv("FILEBASE_S3_REGION")),
		config.WithEndpointResolver(
			aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:           os.Getenv("FILEBASE_S3_ENDPOINT"),
					SigningRegion: os.Getenv("FILEBASE_S3_REGION"),
				}, nil
			}),
		),
	)

	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true // obrigatório no Filebase
	})

	return client, nil
}
