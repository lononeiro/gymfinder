package utils

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"os"
	"strings"
	"time"

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

// Cria sessão Filebase
func NewFilebaseSession() (*session.Session, error) {

	access := strings.TrimSpace(os.Getenv("FILEBASE_ACCESS_KEY"))
	secret := strings.TrimSpace(os.Getenv("FILEBASE_SECRET_KEY"))
	endpoint := strings.TrimSpace(os.Getenv("FILEBASE_ENDPOINT"))

	if access == "" || secret == "" || endpoint == "" {
		return nil, fmt.Errorf("variáveis de ambiente faltando: FILEBASE_ACCESS_KEY, FILEBASE_SECRET_KEY, FILEBASE_ENDPOINT")
	}

	cfg := &aws.Config{
		Endpoint:         aws.String(endpoint),
		Region:           aws.String("us-east-1"),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials(access, secret, ""),
	}

	sess, err := session.NewSession(cfg)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar sessão filebase: %w", err)
	}

	return sess, nil
}

// Faz upload e retorna URL pública do IPFS
func UploadToFilebase(file multipart.File, filename string) (string, error) {
	defer file.Close()

	// Lê arquivo para buffer
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(file)
	if err != nil {
		return "", fmt.Errorf("erro ao ler arquivo: %w", err)
	}

	// Sessão Filebase
	sess, err := NewFilebaseSession()
	if err != nil {
		return "", err
	}
	client := s3.New(sess)

	bucket := os.Getenv("FILEBASE_BUCKET")
	if bucket == "" {
		return "", fmt.Errorf("FILEBASE_BUCKET não definido")
	}

	contentType := getContentType(filename)

	// Upload
	_, err = client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(buf.Bytes()),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", fmt.Errorf("erro no upload: %w", err)
	}

	// ===== Buscar CID no metadata =====

	var head *s3.HeadObjectOutput
	cid := ""

	for attempt := 1; attempt <= 15; attempt++ {

		head, err = client.HeadObject(&s3.HeadObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(filename),
		})
		if err != nil {
			time.Sleep(300 * time.Millisecond)
			continue
		}

		// Filebase usa "Ipfs-Hash"
		for k, v := range head.Metadata {
			if v == nil {
				continue
			}

			key := strings.ToLower(k)

			if key == "ipfs-hash" || strings.Contains(key, "ipfs") {
				if *v != "" {
					cid = *v
					break
				}
			}
		}

		if cid != "" {
			break
		}

		time.Sleep(400 * time.Millisecond)
	}

	if cid == "" {
		return "", fmt.Errorf("CID não encontrado no metadata: %+v", head.Metadata)
	}

	// MONTA URL com SEU gateway Filebase
	url := fmt.Sprintf("https://future-coffee-galliform.myfilebase.com/ipfs/%s", cid)

	return url, nil
}

func getContentType(filename string) string {
	lower := strings.ToLower(filename)
	switch {
	case strings.HasSuffix(lower, ".jpg"), strings.HasSuffix(lower, ".jpeg"):
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

func TestFilebaseConnection() error {
	sess, err := NewFilebaseSession()
	if err != nil {
		return err
	}

	client := s3.New(sess)

	bucket := os.Getenv("FILEBASE_BUCKET")
	if bucket == "" {
		return fmt.Errorf("FILEBASE_BUCKET não definido")
	}

	fmt.Println("🔍 Testando acesso ao Filebase...")

	_, err = client.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return fmt.Errorf("erro ao listar buckets: %w", err)
	}

	fmt.Println("✅ Conexão OK")

	return nil
}
