package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// Tenta carregar .env automaticamente (silencioso)
func init() {
	_ = godotenv.Load()
}

// verifica e retorna as variáveis necessárias (ou erro listando as faltantes)
func mustGetFilebaseEnv() (accessKey, secretKey, region, endpoint, bucket string, err error) {
	accessKey = strings.TrimSpace(os.Getenv("FILEBASE_S3_ACCESS_KEY"))
	secretKey = strings.TrimSpace(os.Getenv("FILEBASE_S3_SECRET_KEY"))
	region = strings.TrimSpace(os.Getenv("FILEBASE_S3_REGION"))
	endpoint = strings.TrimSpace(os.Getenv("FILEBASE_S3_ENDPOINT"))
	fmt.Println(accessKey, secretKey, region, endpoint)

	// aceitar FILEBASE_BUCKET (mais usado) ou FILEBASE_S3_BUCKET (alternativa)
	bucket = strings.TrimSpace(os.Getenv("FILEBASE_BUCKET"))
	if bucket == "" {
		bucket = strings.TrimSpace(os.Getenv("FILEBASE_S3_BUCKET"))
	}

	var missing []string
	if accessKey == "" {
		missing = append(missing, "FILEBASE_S3_ACCESS_KEY")
	}
	if secretKey == "" {
		missing = append(missing, "FILEBASE_S3_SECRET_KEY")
	}
	if region == "" {
		missing = append(missing, "FILEBASE_S3_REGION")
	}
	if endpoint == "" {
		missing = append(missing, "FILEBASE_S3_ENDPOINT")
	}
	if bucket == "" {
		missing = append(missing, "FILEBASE_BUCKET or FILEBASE_S3_BUCKET")
	}

	if len(missing) > 0 {
		err = fmt.Errorf("variáveis de ambiente faltando: %s (carregue .env ou defina no ambiente)", strings.Join(missing, ", "))
	}
	return
}

// NewFilebaseClient cria o cliente S3 (Filebase) usando SDK v2.
// Retorna erro claro se alguma variável estiver faltando.
func NewFilebaseClient(ctx context.Context) (*s3.Client, error) {
	access, secret, region, endpoint, _, err := mustGetFilebaseEnv()
	if err != nil {
		return nil, fmt.Errorf("configuração Filebase inválida: %w", err)
	}

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(access, secret, "")),
		// usar WithEndpointResolverWithOptions para aceitar o tipo com options
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{
						URL:           endpoint,
						SigningRegion: region,
					}, nil
				},
			),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar config AWS: %w", err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	return client, nil
}

// UploadToFilebase faz upload de um multipart.File para o bucket configurado.
// Retorna a URL pública do objeto ou erro detalhado.
func UploadToFilebase(file multipart.File, filename string) (string, error) {
	// lê tudo (se quiser evitar OOM para arquivos gigantes, trocar para multipart upload)
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return "", fmt.Errorf("erro ao ler arquivo do request: %w", err)
	}

	ctx := context.Background()
	client, err := NewFilebaseClient(ctx)
	if err != nil {
		return "", fmt.Errorf("erro ao criar cliente Filebase: %w", err)
	}

	// obtém bucket (pode vir de FILEBASE_BUCKET ou FILEBASE_S3_BUCKET)
	bucket := strings.TrimSpace(os.Getenv("FILEBASE_BUCKET"))
	if bucket == "" {
		bucket = strings.TrimSpace(os.Getenv("FILEBASE_S3_BUCKET"))
	}
	if bucket == "" {
		return "", fmt.Errorf("bucket não configurado: defina FILEBASE_BUCKET ou FILEBASE_S3_BUCKET")
	}

	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &filename,
		Body:   bytes.NewReader(buf.Bytes()),
		ACL:    s3types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		return "", fmt.Errorf("erro ao fazer upload para Filebase: %w", err)
	}

	publicURL := fmt.Sprintf("https://%s.s3.filebase.com/%s", bucket, filename)
	return publicURL, nil
}
