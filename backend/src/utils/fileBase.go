package utils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

// Carrega .env automaticamente
func init() {
	_ = godotenv.Load()
}

func getFilebaseConfig() (accessKey, secretKey, region, endpoint, bucket string, err error) {
	accessKey = strings.TrimSpace(os.Getenv("FILEBASE_S3_ACCESS_KEY"))
	secretKey = strings.TrimSpace(os.Getenv("FILEBASE_S3_SECRET_KEY"))
	region = strings.TrimSpace(os.Getenv("FILEBASE_S3_REGION"))
	endpoint = strings.TrimSpace(os.Getenv("FILEBASE_S3_ENDPOINT"))
	bucket = strings.TrimSpace(os.Getenv("FILEBASE_BUCKET"))

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
		missing = append(missing, "FILEBASE_BUCKET")
	}

	if len(missing) > 0 {
		err = fmt.Errorf("variáveis de ambiente faltando: %s", strings.Join(missing, ", "))
	}

	return
}

// NewFilebaseSession cria uma sessão S3 compatível com Filebase usando AWS SDK v1
func NewFilebaseSession() (*session.Session, error) {
	access, secret, region, endpoint, _, err := getFilebaseConfig()
	if err != nil {
		return nil, fmt.Errorf("configuração Filebase inválida: %w", err)
	}

	// Configuração EXATA como na documentação Filebase
	s3Config := aws.Config{
		Credentials:      credentials.NewStaticCredentials(access, secret, ""),
		Endpoint:         aws.String(endpoint),
		Region:           aws.String(region),
		S3ForcePathStyle: aws.Bool(true), // ← ESSENCIAL para Filebase
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  s3Config,
		Profile: "filebase",
	})
	if err != nil {
		return nil, fmt.Errorf("erro ao criar sessão Filebase: %w", err)
	}

	return sess, nil
}

// UploadToFilebase faz upload de um multipart.File para o Filebase
func UploadToFilebase(file multipart.File, filename string) (string, error) {
	// Lê o conteúdo do arquivo para bytes
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return "", fmt.Errorf("erro ao ler arquivo: %w", err)
	}

	// Cria sessão
	sess, err := NewFilebaseSession()
	if err != nil {
		return "", fmt.Errorf("erro ao criar sessão: %w", err)
	}

	// Cria cliente S3
	s3Client := s3.New(sess)

	// Obtém bucket
	_, _, _, _, bucket, err := getFilebaseConfig()
	if err != nil {
		return "", fmt.Errorf("erro ao obter configurações: %w", err)
	}

	// Determina content type
	contentType := getContentType(filename)

	// Faz upload (igual à documentação)
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(buf.Bytes()),
		ContentType: aws.String(contentType),
	})

	if err != nil {
		return "", fmt.Errorf("erro ao fazer upload para Filebase: %w", err)
	}

	// URL pública
	publicURL := fmt.Sprintf("https://s3.filebase.com/%s/%s", bucket, filename)
	return publicURL, nil
}

// getContentType determina o content type baseado na extensão
func getContentType(filename string) string {
	lower := strings.ToLower(filename)
	switch {
	case strings.HasSuffix(lower, ".jpg") || strings.HasSuffix(lower, ".jpeg"):
		return "image/jpeg"
	case strings.HasSuffix(lower, ".png"):
		return "image/png"
	case strings.HasSuffix(lower, ".gif"):
		return "image/gif"
	case strings.HasSuffix(lower, ".webp"):
		return "image/webp"
	default:
		return "application/octet-stream"
	}
}

// TestFilebaseConnection testa a conexão com Filebase usando SDK v1
func TestFilebaseConnection() error {
	sess, err := NewFilebaseSession()
	if err != nil {
		return fmt.Errorf("falha ao criar sessão: %w", err)
	}

	s3Client := s3.New(sess)

	_, _, _, _, bucket, err := getFilebaseConfig()
	if err != nil {
		return fmt.Errorf("falha ao obter configurações: %w", err)
	}

	fmt.Printf("🔍 Testando conexão com Filebase (SDK v1)...\n")
	fmt.Printf("📦 Bucket: %s\n", bucket)

	// Lista buckets para testar permissões
	result, err := s3Client.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return fmt.Errorf("❌ Falha ao listar buckets: %w", err)
	}

	fmt.Printf("✅ Conexão básica OK - %d buckets encontrados:\n", len(result.Buckets))
	for _, b := range result.Buckets {
		fmt.Printf("   - %s\n", aws.StringValue(b.Name))
	}

	// Verifica se o bucket existe
	_, err = s3Client.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return fmt.Errorf("❌ Bucket '%s' não existe ou não está acessível: %w", bucket, err)
	}

	fmt.Printf("✅ Bucket '%s' está acessível\n", bucket)

	// Testa permissões de escrita
	testKey := "test-permission.txt"
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(testKey),
		Body:   strings.NewReader("teste de permissão"),
	})

	if err != nil {
		return fmt.Errorf("❌ Sem permissão de escrita no bucket: %w", err)
	}

	// Limpa o arquivo de teste
	_, err = s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(testKey),
	})

	fmt.Printf("✅ Permissões de escrita OK\n")
	fmt.Printf("🎉 Filebase configurado corretamente com SDK v1!\n")

	return nil
}
