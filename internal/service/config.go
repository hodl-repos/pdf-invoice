package service

type Config struct {
	Port string `env:"PORT, default=12003"`
}
