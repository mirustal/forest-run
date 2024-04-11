package boot

import (
	"log"

	"github.com/spf13/viper"
)

const (
	DevEnv  = "dev"
	ProdEnv = "prod"
)

type Env struct {
	AppEnv            string `mapstructure:"APP_ENV"`
	ServerAddress     string `mapstructure:"SERVER_ADDRESS"`
	LogSaveFile       string `mapstructure:"LOG_SAVE_FILE"`
	MaxLogFileSize    int    `mapstructure:"MAX_LOG_FILE_SIZE_MB"`
	MaxLogFileBackups int    `mapstructure:"MAX_LOG_FILE_BACKUP"`
	MaxLogFileAge     int    `mapstructure:"MAX_LOG_FILE_AGE_DAYS"`
	DBUrl             string `mapstructure:"DB_URL"`
	SecureKey         string `mapstructure:"SECURE_KEY"`
}

func NewEnv() (*Env, error) {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
		return nil, err
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
		return nil, err
	}

	if env.AppEnv == DevEnv {
		log.Println("The App is running in development env")
	}

	return &env, nil
}
