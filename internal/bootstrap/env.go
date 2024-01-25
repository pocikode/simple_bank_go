package bootstrap

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	AppEnv  string `mapstructure:"APP_ENV"`
	AppPort string `mapstructure:"APP_PORT"`
	DBHost  string `mapstructure:"DB_HOST"`
	DBPort  string `mapstructure:"DB_PORT"`
	DBUser  string `mapstructure:"DB_USER"`
	DBPass  string `mapstructure:"DB_PASS"`
	DBName  string `mapstructure:"DB_NAME"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Can't find the file .env :", err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal("Environment can't be loaded:", err)
	}

	if env.AppEnv == "development" {
		log.Println("This App is running in development env")
	}

	return &env
}
