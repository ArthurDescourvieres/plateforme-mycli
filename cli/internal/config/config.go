package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	URL       string
	AccessKey string
	SecretKey string
	Region    string
}

func Load() Config {
	for _, path := range []string{"../docker/.env", "docker/.env"} {
		if err := godotenv.Load(path); err == nil {
			break
		}
	}

	return Config{
		URL:       os.Getenv("MINIO_ENDPOINT"),
		AccessKey: os.Getenv("MINIO_USER"),
		SecretKey: os.Getenv("MINIO_PASS"),
		Region:    os.Getenv("MINIO_REGION"),
	}
}
