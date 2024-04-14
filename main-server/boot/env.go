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
	AppEnv        string `mapstructure:"APP_ENV"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	LoggerConfig
	JWTConfig
	DBConfig
}

type LoggerConfig struct {
	LogSaveFile       string `mapstructure:"LOG_SAVE_FILE"`
	MaxLogFileSize    int    `mapstructure:"MAX_LOG_FILE_SIZE_MB"`
	MaxLogFileBackups int    `mapstructure:"MAX_LOG_FILE_BACKUP"`
	MaxLogFileAge     int    `mapstructure:"MAX_LOG_FILE_AGE_DAYS"`
}

type DBConfig struct {
	DBUrl string `mapstructure:"DB_URL"`
}

type JWTConfig struct {
	SecureKey            string        `mapstructure:"SECURE_KEY"`
	JWTTokenLifeTime     time.Duration `mapstructure:"JWT_TOKEN_LIFE_TIME"`
	RefreshTokenLifeTime time.Duration `mapstructure:"REFRESH_TOKEN_LIFE_TIME"`
}

func NewEnv() (env Env, err error) {
	viper.SetConfigFile(".env")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
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
