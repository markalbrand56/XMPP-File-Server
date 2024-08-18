package configs

import (
	"XMPP-File-Server/internal/database"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

const (
	MAX_UPLOAD_SIZE = 1024 * 1024 * 6 // 6MB
)

// connectS3 se conecta a AWS S3
func connectS3() {
	envs, err := LoadAmazonCredentials()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new S3 client
	awsSession, err := session.NewSession(&aws.Config{
		Region:      aws.String(envs["AWS_REGION"]),
		Credentials: credentials.NewStaticCredentials(envs["AWS_ACCESS"], envs["AWS_SECRET"], ""),
	})

	if err != nil {
		log.Fatal(err)
	}

	// Create S3 service client
	newSession := s3.New(awsSession)
	database.Instance = database.NewBucket(newSession, envs["AWS_BUCKET"])

	log.Println("Connected to S3")
}

func init() {
	connectS3()
}
