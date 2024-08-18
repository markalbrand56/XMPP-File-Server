package database

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"mime/multipart"
)

var Instance *Bucket

type Bucket struct {
	S3   *s3.S3
	Name string
}

// NewBucket crea una nueva instancia de Bucket
func NewBucket(s3 *s3.S3, name string) *Bucket {
	return &Bucket{
		S3:   s3,
		Name: name,
	}
}

// Insert inserta un archivo en el bucket de S3
func (b *Bucket) Insert(path string, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}

	buffer := make([]byte, file.Size)
	_, err = src.Read(buffer)

	if err != nil {
		return err
	}

	defer src.Close()

	_, err = b.S3.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(b.Name),
		Key:                aws.String(path),
		Body:               bytes.NewReader(buffer),
		ContentType:        aws.String("application/octet-stream"),
		ContentDisposition: aws.String("attachment"),
		ContentLength:      aws.Int64(file.Size),
	})

	if err != nil {
		return err
	}

	return nil
}

// GetFile obtiene un archivo del bucket de S3
func (b *Bucket) GetFile(path string) (*bytes.Buffer, error) {
	result, err := b.S3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(b.Name),
		Key:    aws.String(path),
	})

	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(result.Body)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

// DeleteFile elimina un archivo del bucket de S3
func (b *Bucket) DeleteFile(path string) error {
	_, err := b.S3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(b.Name),
		Key:    aws.String(path),
	})

	if err != nil {
		return err
	}

	return nil
}

// DeleteFolder elimina una carpeta del bucket de S3
func (b *Bucket) DeleteFolder(folderPath string) error {
	// List all objects in the folder
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(b.Name),
		Prefix: aws.String(folderPath),
	}
	objects, err := b.S3.ListObjectsV2(input)
	if err != nil {
		return err
	}

	// Delete each object in the folder
	for _, object := range objects.Contents {
		deleteInput := &s3.DeleteObjectInput{
			Bucket: aws.String(b.Name),
			Key:    object.Key,
		}
		_, err := b.S3.DeleteObject(deleteInput)
		if err != nil {
			return err
		}
	}

	return nil
}
