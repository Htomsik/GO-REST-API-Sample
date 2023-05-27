package apiServer

// Config ...
type Config struct {
	Addr     string `toml:"addr"`
	LogLevel string `toml:"logLevel"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Addr:     ":3030",
		LogLevel: "debug",
	}
}
