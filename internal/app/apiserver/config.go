package apiserver

// Config - description of config file
type Config struct {
	Host     string `toml:"host"`
	Loglevel string `toml:"log_level"`
}

// NewConfig - return config
func NewConfig() *Config {
	return &Config{
		Host:     ":8888",
		Loglevel: "info",
	}
}
