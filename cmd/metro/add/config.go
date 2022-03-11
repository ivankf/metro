package add

import "os"

type Config struct {
	author   string `toml:"author"`
	template string
	body     string
}

func NewConfig() *Config {
	return &Config{
		author:   os.Getenv("METRO_AUTHOR"),
		template: os.Getenv("METRO_TEMPLATE"),
		body:     os.Getenv("METRO_BODY"),
	}
}
