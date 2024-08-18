package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	environmentVariables = []string{
		"AWS_ACCESS_KEY",
		"AWS_SECRET",
		"AWS_REGION",
		"AWS_BUCKET",
		"URL",
	}
)

// LoadAmazonCredentials carga las credenciales de AWS de las variables de entorno
func LoadAmazonCredentials() (map[string]string, error) {
	envs := make(map[string]string)

	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, loading from system environment")
	}

	for _, env := range environmentVariables {
		envs[env] = os.Getenv(env)

		if envs[env] == "" {
			return nil, fmt.Errorf("credentials (%s) not found", env)
		}
	}

	return envs, nil
}
