package utils

import (
	"bytes"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

const BucketName = "gymfinder" // o nome do bucket que você criou no filebase

func UploadToFilebase(file multipart.File, filename string) (string, error) {
	sizeBuf := new(bytes.Buffer)
	sizeBuf.ReadFrom(file)
	fileBytes := sizeBuf.Bytes()

	svc := NewFilebaseClient()

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(fileBytes),
		ACL:    aws.String("public-read"),
	})

	if err != nil {
		return "", err
	}

	// URL pública do Filebase
	url := "https://" + BucketName + ".s3.filebase.com/" + filename
	return url, nil
}
