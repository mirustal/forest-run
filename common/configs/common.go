package configs

import "errors"

type EnvType string

const (
	DevEnv  = EnvType("dev")
	ProdEnv = EnvType("prod")
)

type CommonConfig struct {
	AppEnv        EnvType `yaml:"appEnv"`
	ServerAddress string  `yaml:"serverAddress"`
}

type LoggerConfig struct {
	LogSaveFile       string `yaml:"logSaveFile"`
	MaxLogFileSize    int    `yaml:"maxLogFileSize"`
	MaxLogFileBackups int    `yaml:"maxLogFileBackups"`
	MaxLogFileAge     int    `yaml:"maxLogFileAge"`
}

func (cfg LoggerConfig) Validate() error {
	if len(cfg.LogSaveFile) == 0 {
		return errors.New("LogSaveFile length = 0")
	}

	if cfg.MaxLogFileSize <= 0 {
		return errors.New("MaxLogFileSize <= 0")
	}

	if cfg.MaxLogFileBackups < 0 {
		return errors.New("MaxLogFileBackups < 0")
	}

	if cfg.MaxLogFileAge <= 0 {
		return errors.New("MaxLogFileAge <= 0")
	}

	return nil
}
