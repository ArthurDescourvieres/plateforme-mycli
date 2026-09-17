package config

type Config struct {
	URL       string
	AccessKey string
	SecretKey string
	Region    string
}

func Load() Config {
	return Config{}
}
