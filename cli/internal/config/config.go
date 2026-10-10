package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	URL       string
	AccessKey string
	SecretKey string
	Region    string
}

const (
	defaultURL    = "http://localhost:9000"
	defaultRegion = "us-east-1"
)

type Profile struct {
	URL       string `json:"url"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Region    string `json:"region"`
}

type fileConfig struct {
	Default string             `json:"default"`
	Aliases map[string]Profile `json:"aliases"`
}

func configPath() (string, error) {
	if path := os.Getenv("MYCLI_CONFIG"); path != "" {
		return path, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".mycli", "config.json"), nil
}

func loadConfigFileFromPath(path string) (fileConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return fileConfig{}, err
	}

	var cfg fileConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return fileConfig{}, err
	}

	return cfg, nil
}

func loadConfigFile() (fileConfig, error) {
	path, err := configPath()
	if err != nil {
		return fileConfig{}, err
	}

	return loadConfigFileFromPath(path)
}

func loadDefaultProfile() (Profile, error) {
	cfg, err := loadConfigFile()
	if err != nil {
		return Profile{}, err
	}

	profile, ok := cfg.Aliases[cfg.Default]
	if !ok {
		return Profile{}, fmt.Errorf("default profile %q not found", cfg.Default)
	}

	return profile, nil
}

func mergeEnv(profile Profile) Profile {
	if value, ok := lookupEnv("MYCLI_URL", "MINIO_ENDPOINT"); ok {
		profile.URL = value
	}

	if value, ok := lookupEnv("MYCLI_ACCESS_KEY", "MINIO_USER"); ok {
		profile.AccessKey = value
	}

	if value, ok := lookupEnv("MYCLI_SECRET_KEY", "MINIO_PASS"); ok {
		profile.SecretKey = value
	}

	if value, ok := lookupEnv("MYCLI_REGION", "MINIO_REGION"); ok {
		profile.Region = value
	}

	return profile
}

func lookupEnv(primary string, fallback string) (string, bool) {
	if value, ok := os.LookupEnv(primary); ok {
		return value, true
	}

	return os.LookupEnv(fallback)
}

func Load() Config {
	for _, path := range []string{"../docker/.env", "docker/.env"} {
		if err := godotenv.Load(path); err == nil {
			break
		}
	}

	profile := Profile{
		URL:    defaultURL,
		Region: defaultRegion,
	}

	if fileProfile, err := loadDefaultProfile(); err == nil {
		profile = fileProfile
	}

	profile = mergeEnv(profile)

	return Config(profile)
}
