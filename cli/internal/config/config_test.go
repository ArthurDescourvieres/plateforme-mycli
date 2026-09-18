package config

import (
	"os"
	"path/filepath"
	"testing"
)

// TestLoadFromEnv vérifie que la configuration peut être chargée depuis les variables d'environnement MYCLI_*
func TestLoadFromEnv(t *testing.T) {
	t.Setenv("MYCLI_URL", "http://serveur-test:9000")
	t.Setenv("MYCLI_ACCESS_KEY", "test-access")
	t.Setenv("MYCLI_SECRET_KEY", "test-secret")
	t.Setenv("MYCLI_REGION", "us-east-1")

	cfg := Load()

	if cfg.URL != "http://serveur-test:9000" {
		t.Errorf("expected URL %q, got %q", "http://serveur-test:9000", cfg.URL)
	}

	if cfg.AccessKey != "test-access" {
		t.Errorf("expected AccessKey %q, got %q", "test-access", cfg.AccessKey)
	}

	if cfg.SecretKey != "test-secret" {
		t.Errorf("expected SecretKey %q, got %q", "test-secret", cfg.SecretKey)
	}

	if cfg.Region != "us-east-1" {
		t.Errorf("expected Region %q, got %q", "us-east-1", cfg.Region)
	}
}

// TestLoadDefaultProfile vérifie que le profil par défaut est correctement récupéré depuis le fichier de configuration.
func TestLoadDefaultProfile(t *testing.T) {
	profile, err := loadDefaultProfile()
	if err != nil {
		t.Fatalf("failed to load default profile: %v", err)
	}

	if profile.URL != "http://localhost:9000" {
		t.Errorf("expected URL %q, got %q", "http://localhost:9000", profile.URL)
	}

	if profile.AccessKey != "test-access" {
		t.Errorf("expected AccessKey %q, got %q", "test-access", profile.AccessKey)
	}

	if profile.SecretKey != "test-secret" {
		t.Errorf("expected SecretKey %q, got %q", "test-secret", profile.SecretKey)
	}

	if profile.Region != "us-east-1" {
		t.Errorf("expected Region %q, got %q", "us-east-1", profile.Region)
	}
}

// TestLoadEnvironmentOverridesFile vérifie que les variables d'environnement ont priorité sur les valeurs du fichier.
func TestLoadEnvironmentOverridesFile(t *testing.T) {
	t.Setenv("MYCLI_URL", "http://serveur-env:9000")

	cfg := Load()

	if cfg.URL != "http://serveur-env:9000" {
		t.Errorf("expected URL %q, got %q", "http://serveur-env:9000", cfg.URL)
	}
}

// TestLoadConfigFile vérifie que le fichier JSON de configuration est correctement lu et que les valeurs du profil sont récupérées.
func TestLoadConfigFile(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.json")

	configContent := `{
		"default": "minio-test",
		"aliases": {
			"minio-test": {
				"url": "http://minio-test:9000",
				"access_key": "file-access",
				"secret_key": "file-secret",
				"region": "eu-west-1"
			}
		}
	}`

	if err := os.WriteFile(configPath, []byte(configContent), 0600); err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	cfg, err := loadConfigFileFromPath(configPath)
	if err != nil {
		t.Fatalf("failed to load test config: %v", err)
	}

	profile := cfg.Aliases[cfg.Default]

	if profile.URL != "http://minio-test:9000" {
		t.Errorf("expected URL %q, got %q", "http://minio-test:9000", profile.URL)
	}

	if profile.AccessKey != "file-access" {
		t.Errorf("expected AccessKey %q, got %q", "file-access", profile.AccessKey)
	}

	if profile.SecretKey != "file-secret" {
		t.Errorf("expected SecretKey %q, got %q", "file-secret", profile.SecretKey)
	}

	if profile.Region != "eu-west-1" {
		t.Errorf("expected Region %q, got %q", "eu-west-1", profile.Region)
	}
}

// TestLoadDefaultValues vérifie que les valeurs par défaut // sont utilisées lorsqu'aucun fichier ni variable d'environnement // ne fournit de configuration.
func TestLoadDefaultValues(t *testing.T) {
	t.Setenv("MYCLI_CONFIG", filepath.Join(t.TempDir(), "config.json"))

	os.Unsetenv("MYCLI_URL")
	os.Unsetenv("MYCLI_ACCESS_KEY")
	os.Unsetenv("MYCLI_SECRET_KEY")
	os.Unsetenv("MYCLI_REGION")

	cfg := Load()

	if cfg.URL != defaultURL {
		t.Errorf("expected default URL %q, got %q", defaultURL, cfg.URL)
	}

	if cfg.AccessKey != "" {
		t.Errorf("expected empty AccessKey, got %q", cfg.AccessKey)
	}

	if cfg.SecretKey != "" {
		t.Errorf("expected empty SecretKey, got %q", cfg.SecretKey)
	}

	if cfg.Region != defaultRegion {
		t.Errorf("expected default Region %q, got %q", defaultRegion, cfg.Region)
	}
}
