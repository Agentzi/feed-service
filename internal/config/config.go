package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port        string
	DatabaseUrl string
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No .env file found, relying on environment variables: %v", err)
	}

	config := &Config{
		Port:        viper.GetString("PORT"),
		DatabaseUrl: viper.GetString("DATABASE_URL"),
	}

	if config.Port == "" {
		config.Port = "5000"
	}

	return config, nil
}
