package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/apiServer"
	"log"
)

var (
	cfgPath string
)

func init() {
	flag.StringVar(&cfgPath, "config-path", "configs/apiServer.toml", "path to cfg file")
}

func main() {
	flag.Parse()

	cfg := apiServer.NewConfig()

	if _, err := toml.DecodeFile(cfgPath, cfg); err != nil {
		log.Fatal(err)
	}

	if err := apiServer.Start(cfg); err != nil {
		log.Fatal(err)
	}
}
