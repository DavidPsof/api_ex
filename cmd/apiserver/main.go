package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/DavidPsof/api_ex/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to configuration file")
}

func main() {
	flag.Parse()
	cfg := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, cfg)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.NewAPIServer(cfg)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
