package utils

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// ================================
// CLIENTE FILEBASE (SDK V2)
// ================================
func NewFilebaseClient(ctx context.Context) (*s3.Client, error) {

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			os.Getenv("FILEBASE_S3_ACCESS_KEY"),
			os.Getenv("FILEBASE_S3_SECRET_KEY"),
			"",
		)),
		config.WithRegion(os.Getenv("FILEBASE_S3_REGION")),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{
						URL:           os.Getenv("FILEBASE_S3_ENDPOINT"),
						SigningRegion: os.Getenv("FILEBASE_S3_REGION"),
					}, nil
				},
			),
		),
	)

	if err != nil {
		return nil, fmt.Errorf("erro ao carregar config do Filebase: %w", err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true // obrigatório
	})

	return client, nil
}

// ================================
// UPLOAD PARA FILEBASE
// ================================
func UploadToFilebase(file multipart.File, filename string) (string, error) {
	ctx := context.TODO()

	// Lê arquivo inteiro
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(file)
	if err != nil {
		return "", fmt.Errorf("erro ao ler arquivo: %w", err)
	}

	client, err := NewFilebaseClient(ctx)
	if err != nil {
		return "", err
	}

	bucket := os.Getenv("FILEBASE_BUCKET")

	// Upload usando SDK v2
	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(buf.Bytes()),
		ACL:    types.ObjectCannedACLPublicRead, // ✔ ACL correta no SDK v2
	})

	if err != nil {
		return "", fmt.Errorf("erro ao fazer upload para Filebase: %w", err)
	}

	// URL pública
	publicURL := fmt.Sprintf("https://%s.s3.filebase.com/%s", bucket, filename)

	return publicURL, nil
}
