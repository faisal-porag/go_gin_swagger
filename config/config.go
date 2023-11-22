package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	ApplicationPort  int  `env:"APPLICATION_PORT" envDefault:"8070"`
	DebugMode        bool `env:"DEBUG_MODE" envDefault:"false"`
	IsProductionMode bool `env:"IS_PRODUCTION_MODE" envDefault:"false"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	return cfg, err
}
