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
	MAX_UPLOAD_SIZE = 1024 * 1024 * 10 // 10MB
)

var URL string

// connectS3 se conecta a AWS S3
func connectS3() {
	envs, err := LoadAmazonCredentials()
	if err != nil {
		log.Fatal(err)
	}

	// For some reason, the AWS SDK doesn't like accessing the environment variables directly from the map
	region := envs["AWS_REGION"]
	accessKey := envs["AWS_ACCESS_KEY"]
	secret := envs["AWS_SECRET"]

	// Create a new S3 client
	awsSession, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secret, ""),
	})

	if err != nil {
		log.Fatal(err)
	}

	// Create S3 service client
	newSession := s3.New(awsSession)
	database.Instance = database.NewBucket(newSession, envs["AWS_BUCKET"])

	log.Println("Connected to S3")

	URL = envs["URL"]
}

func init() {
	connectS3()
}
