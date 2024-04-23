package boot

import (
	"errors"
	"forest-run/common/configs"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Env struct {
	configs.CommonConfig `yaml:"commonConfig"`
	configs.LoggerConfig `yaml:"loggerConfig"`
	JWTConfig            `yaml:"JWTConfig"`
	DBConfig             `yaml:"DBConfig"`
}

func (e Env) validate() error {
	if err := e.JWTConfig.validate(); err != nil {
		return err
	}

	if err := e.LoggerConfig.Validate(); err != nil {
		return err
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

	if env.AppEnv == configs.DevEnv {
		log.Println("The App is running in development env")
	}

	return env, env.validate()
}
