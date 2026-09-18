package config

import "os"

type Config struct {
	URL       string
	AccessKey string
	SecretKey string
	Region    string
}

func Load() Config {
	return Config{
		URL:       os.Getenv("MINIO_ENDPOINT"),
		AccessKey: os.Getenv("MINIO_USER"),
		SecretKey: os.Getenv("MINIO_PASS"),
		Region:    os.Getenv("MINIO_REGION"),
	}
}
