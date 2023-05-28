package store

type Config struct {
	DatabaseURL  string `toml:"databaseURL"`
	DatabaseType string `toml:"databaseType"`
}

func NewConfig() *Config {
	return &Config{}
}
