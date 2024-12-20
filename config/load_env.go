package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBUrl string `mapstructure:"DB_URL"`
}

func LoadEnv() (Config, error) {
	var config Config

	// Automatic read from environment
	viper.AutomaticEnv()

	// Optional read from file
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No env file found")
	}

	// Unmarshal the configuration into the Config struct
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	// Validate required configuration
	if config.DBUrl == "" {
		return config, fmt.Errorf("DB_URL is required")
	}

	return config, nil
}
