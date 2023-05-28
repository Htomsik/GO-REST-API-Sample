package apiServer

import "github.com/Htomsik/GO-REST-API-Sample/internal/app/store"

// Config ...
type Config struct {
	Addr     string `toml:"addr"`
	LogLevel string `toml:"logLevel"`
	Store    *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Addr:     ":3030",
		LogLevel: "debug",
	}
}
