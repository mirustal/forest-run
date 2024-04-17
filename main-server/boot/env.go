package boot

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

const (
	DevEnv  = "dev"
	ProdEnv = "prod"
)

type Env struct {
	AppEnv        string `yaml:"appEnv"`
	ServerAddress string `yaml:"serverAddress"`
	DefsPath      string `yaml:"defsPath"` // replace with url to cdn
	LoggerConfig  `yaml:"loggerConfig"`
	JWTConfig     `yaml:"JWTConfig"`
	DBConfig      `yaml:"DBConfig"`
}

func (e Env) validate() error {
	if err := e.JWTConfig.validate(); err != nil {
		return err
	}

	if err := e.LoggerConfig.validate(); err != nil {
		return err
	}

	return nil
}

type LoggerConfig struct {
	LogSaveFile       string `yaml:"logSaveFile"`
	MaxLogFileSize    int    `yaml:"maxLogFileSize"`
	MaxLogFileBackups int    `yaml:"maxLogFileBackups"`
	MaxLogFileAge     int    `yaml:"maxLogFileAge"`
}

func (cfg LoggerConfig) validate() error {
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

type DBConfig struct {
	DBUrl string `yaml:"DBUrl"`
}

type JWTConfig struct {
	SecureKey            string        `yaml:"secureKey"`
	JWTTokenLifeTime     time.Duration `yaml:"JWTTokenLifeTime"`
	RefreshTokenLifeTime time.Duration `yaml:"refreshTokenLifeTime"`
}

func (cfg JWTConfig) validate() error {
	if cfg.JWTTokenLifeTime <= 0 {
		return errors.New("JWTTokenLifeTime <= 0")
	}

	if len(cfg.SecureKey) == 0 {
		return errors.New("length of secure key is 0")
	}

	if cfg.RefreshTokenLifeTime <= 0 {
		return errors.New("RefreshTokenLifeTime <= 0")
	}

	return nil
}

func NewEnv() (env Env, err error) {
	viper.SetConfigFile("conf.yaml")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the config file : ", err)
		return env, err
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
		return env, err
	}

	if env.AppEnv == DevEnv {
		log.Println("The App is running in development env")
	}

	return env, env.validate()
}
