package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config repr├®sente la configuration finale utilis├®e par l'application.
type Config struct {
	URL       string
	AccessKey string
	SecretKey string
	Region    string
}

// Valeurs utilis├®es lorsque aucune configuration n'est fournie // par les variables d'environnement ou le fichier de configuration.
const (
	defaultURL    = "http://localhost:9000"
	defaultRegion = "us-east-1"
)

// Profile repr├®sente un profil de connexion enregistr├® // dans le fichier de configuration JSON.
type Profile struct {
	URL       string `json:"url"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Region    string `json:"region"`
}

// fileConfig repr├®sente la structure compl├¿te du fichier config.json. // Le fichier peut contenir plusieurs profils appel├®s "aliases".
type fileConfig struct {
	Default string             `json:"default"`
	Aliases map[string]Profile `json:"aliases"`
}

// configPath d├®termine le chemin du fichier de configuration. // La variable MYCLI_CONFIG permet de d├®finir un chemin personnalis├®, // principalement utile pour les tests. // En utilisation normale, le fichier se trouve dans ~/.mycli/config.json.
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

// loadConfigFileFromPath lit un fichier JSON ├á partir du chemin fourni // et le transforme en structure fileConfig.
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

// loadConfigFile r├®cup├¿re le chemin du fichier de configuration, // puis charge et d├®code son contenu JSON.
func loadConfigFile() (fileConfig, error) {
	path, err := configPath()
	if err != nil {
		return fileConfig{}, err
	}

	return loadConfigFileFromPath(path)
}

// loadDefaultProfile charge le fichier de configuration // et r├®cup├¿re le profil indiqu├® par la propri├®t├® "default".
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

// mergeEnv applique les variables d'environnement MYCLI_* // sur le profil fourni. // Les variables pr├®sentes ont priorit├® sur les valeurs du fichier.
func mergeEnv(profile Profile) Profile {
	if value, ok := os.LookupEnv("MYCLI_URL"); ok {
		profile.URL = value
	}

	if value, ok := os.LookupEnv("MYCLI_ACCESS_KEY"); ok {
		profile.AccessKey = value
	}

	if value, ok := os.LookupEnv("MYCLI_SECRET_KEY"); ok {
		profile.SecretKey = value
	}

	if value, ok := os.LookupEnv("MYCLI_REGION"); ok {
		profile.Region = value
	}

	return profile
}

// Load construit la configuration finale de MyCLI. // La priorit├® appliqu├®e est : // environnement > fichier de configuration > valeurs par d├®faut.
func Load() Config {
	profile := Profile{
		URL:    defaultURL,
		Region: defaultRegion,
	}

	if fileProfile, err := loadDefaultProfile(); err == nil {
		profile = fileProfile
	}

	profile = mergeEnv(profile)

	return Config{
		URL:       profile.URL,
		AccessKey: profile.AccessKey,
		SecretKey: profile.SecretKey,
		Region:    profile.Region,
	}
}
