package boot

import (
	"forest-run/common/configs"
	"forest-run/common/jwt"
	"forest-run/common/runs"
	"forest-run/main-server/domain"
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	configs.CommonConfig  `yaml:"commonConfig"`
	configs.LoggerConfig  `yaml:"loggerConfig"`
	jwt.JWTConfig         `yaml:"JWTConfig"`
	DBConfig              `yaml:"DBConfig"`
	RealTimeServersConfig `yaml:"RealTimeServers" mapstructure:"RealTimeServers"`
}

func (e Env) validate() error {
	if err := e.JWTConfig.Validate(); err != nil {
		return err
	}

	if err := e.LoggerConfig.Validate(); err != nil {
		return err
	}

	return nil
}

type RealTimeServersConfig map[domain.RegionId]RealTimeServerConfig
type RealTimeServerConfig struct {
	Address             string     `yaml:"Address"`
	ApproximatePosition runs.Point `yaml:"ApproximatePosition"`
}

type DBConfig struct {
	DBUrl string `yaml:"DBUrl"`
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
