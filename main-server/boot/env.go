package boot

import (
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
	LoggerConfig  `yaml:"loggerConfig"`
	JWTConfig     `yaml:"JWTConfig"`
	DBConfig      `yaml:"DBConfig"`
}

type LoggerConfig struct {
	LogSaveFile       string `yaml:"logSaveFile"`
	MaxLogFileSize    int    `yaml:"maxLogFileSize"`
	MaxLogFileBackups int    `yaml:"maxLogFileBackups"`
	MaxLogFileAge     int    `yaml:"maxLogFileAge"`
}

type DBConfig struct {
	DBUrl string `yaml:"DBUrl"`
}

type JWTConfig struct {
	SecureKey            string        `yaml:"secureKey"`
	JWTTokenLifeTime     time.Duration `yaml:"JWTTokenLifeTime"`
	RefreshTokenLifeTime time.Duration `yaml:"refreshTokenLifeTime"`
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

	return env, nil
}
