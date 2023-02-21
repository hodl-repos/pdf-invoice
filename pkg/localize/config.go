package localize

type Config struct {
	//the language keys separated by comma, all small
	LangKeys string `env:"LANGUAGE_KEYS"`
}

func (c *Config) LocalizeServiceConfig() *Config {
	return c
}
