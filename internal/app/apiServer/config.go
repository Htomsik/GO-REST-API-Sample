package apiServer

// Config ...
type Config struct {
	Addr         string `toml:"addr"`
	LogLevel     string `toml:"logLevel"`
	DatabaseType string `toml:"databaseType"`
	DatabaseURL  string `toml:"databaseURL"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Addr:     ":3030",
		LogLevel: "debug",
	}
}
