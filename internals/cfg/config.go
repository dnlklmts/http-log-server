package cfg

import "os"

type Config struct {
	DefaultPath  string
	ProvidedPath string
}

func NewConfig() *Config {
	cfg := new(Config)

	cfg.DefaultPath = "./log.txt"
	if path, ok := os.LookupEnv("APP_LOGFILE_PATH"); ok {
		cfg.ProvidedPath = path
	}

	return cfg
}
