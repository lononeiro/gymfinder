package utils

import (
	"bytes"
	// "context"
	"fmt"
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

// NewFilebaseSession cria e retorna uma sessão da AWS configurada para o Filebase.
func NewFilebaseSession() (*session.Session, error) {
	accessKey := os.Getenv("FILEBASE_ACCESS_KEY")
	secretKey := os.Getenv("FILEBASE_SECRET_KEY")
	endpoint := os.Getenv("FILEBASE_ENDPOINT")

	if accessKey == "" || secretKey == "" || endpoint == "" {
		return nil, fmt.Errorf("FILEBASE_ACCESS_KEY, FILEBASE_SECRET_KEY ou FILEBASE_ENDPOINT não definidos")
	}

	s3Config := &aws.Config{
		Endpoint:         aws.String(endpoint),
		Region:           aws.String("us-east-1"),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
	}

	sess, err := session.NewSession(s3Config)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar sessão Filebase: %w", err)
	}

	return sess, nil
}

// UploadToFilebase faz o upload de um arquivo para o Filebase e retorna a URL pública do IPFS.
func UploadToFilebase(file multipart.File, filename string) (string, error) {
	defer file.Close()
	// ctx := context.TODO()

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(file)
	if err != nil {
		return "", fmt.Errorf("erro ao ler arquivo: %w", err)
	}

	sess, err := NewFilebaseSession()
	if err != nil {
		return "", err
	}

	client := s3.New(sess)

	bucket := os.Getenv("FILEBASE_BUCKET")
	if bucket == "" {
		return "", fmt.Errorf("FILEBASE_BUCKET não definido")
	}

	_, err = client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(buf.Bytes()),
	})
	if err != nil {
		return "", fmt.Errorf("erro no upload: %w", err)
	}

	head, err := client.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return "", fmt.Errorf("erro ao obter metadata: %w", err)
	}

	cid := ""
	if head.Metadata != nil {
		// Verifique se a chave existe e obtenha o valor (que é um *string)
		valuePtr, ok := head.Metadata["x-filebase-object-cid"]

		// Se a chave existir E o ponteiro não for nulo, desreferencie-o
		if ok && valuePtr != nil {
			// Solução: Desreferencie o ponteiro para *string
			cid = *valuePtr
		}
	}

	if cid == "" {
		return "", fmt.Errorf("CID não encontrado no metadado")
	}

	publicURL := fmt.Sprintf("https://ipfs.filebase.io/ipfs/%s", cid)

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
