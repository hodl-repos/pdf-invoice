package config

// Config represents the configuration and associated environment variables for
// the wms components.
type Config struct {
	Port string `env:"PORT, default=12003"`
}
