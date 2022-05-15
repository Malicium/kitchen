package config

import (
	"os"
)

type Config struct {
	HttpAddress  string
	TrustedProxy string
}

func Get() Config {
	return Config{
		HttpAddress:  os.Getenv("HttpAddress"),
		TrustedProxy: os.Getenv("TrustedProxy"),
	}
}
