package service

import "github.com/hodl-repos/pdf-invoice/pkg/localize"

type Config struct {
	LocalizeConfig *localize.Config
	Port           string `env:"PORT, default=12003"`
}

func (c *Config) LocalizeServiceConfig() *localize.Config {
	return c.LocalizeConfig
}
