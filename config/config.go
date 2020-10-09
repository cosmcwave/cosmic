package config

import (
	"context"
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	ServerAddress string
	Registry      string
	TLS           bool
}

func Load(ctx context.Context, confPath string) (*Config, error) {
	var conf *Config
	if _, err := toml.DecodeFile(confPath, &conf); err != nil {
		return nil, err
	}
	log.Println("==> backend config file loaded successfully")
	return conf, nil
}
