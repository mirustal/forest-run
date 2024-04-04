package boot

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv            string `mapstructure:"APP_ENV"`
	ServerAddress     string `mapstructure:"SERVER_ADDRESS"`
	LogSaveFile       string `mapstructure:"LOG_SAVE_FILE"`
	MaxLogFileSize    int    `mapstructure:"MAX_LOG_FILE_SIZE_MB"`
	MaxLogFileBackups int    `mapstructure:"MAX_LOG_FILE_BACKUP"`
	MaxLogFileAge     int    `mapstructure:"MAX_LOG_FILE_AGE_DAYS"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
