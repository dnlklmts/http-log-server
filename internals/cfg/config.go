package cfg

import "os"

type Config struct {
	Port        string
	LogFilePath string
}

func NewConfig() *Config {
	cfg := new(Config)
	cfg.SetUpEnv()

	cfg.Port = os.Getenv("APP_DEFAULT_PORT")
	cfg.LogFilePath = os.Getenv("APP_DEFAULT_LOGFILE")

	if _, ok := os.LookupEnv("APP_LOGFILE_PATH"); ok {
		cfg.LogFilePath = os.Getenv("APP_LOGFILE_PATH")
	}

	return cfg
}

func (cfg *Config) SetUpEnv() {
	os.Setenv("APP_DEFAULT_PORT", "8080")
	os.Setenv("APP_DEFAULT_LOGFILE", "./log.txt")
}
