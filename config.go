package main

import (
	"github.com/caarlos0/env"
	"github.com/caarlos0/env/parsers"
)

type Config struct {
	HttpAddr    string `env:"HTTP_ADDR" envDefault:"0.0.0.0:80"`
	HttpTimeout int    `env:"HTTP_TIMEOUT" envDefault:"5"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	if err := env.ParseWithFuncs(&cfg, env.CustomParsers{
		parsers.URLType: parsers.URLFunc,
	}); err != nil {
		return nil, err
	}

	return &cfg, nil
}
