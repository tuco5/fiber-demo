package config

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/lpernett/godotenv"
)

type Config struct {
	DbUrl string `mapstructure:"DB_URL" validate:"required"`
	Port  string `mapstructure:"PORT"`
	Env   string `mapstructure:"ENV"`
}

func Load(validator *validator.Validate) *Config {
	// Default values
	port := ":3000"
	env := "dev"

	// Load env variables
	godotenv.Load()

	// Check if env variables are set
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	if os.Getenv("ENV") != "" {
		env = os.Getenv("ENV")
	}

	// Convert to struct
	cfg := &Config{
		Port:  port,
		Env:   env,
		DbUrl: os.Getenv("DB_URL"),
	}

	// Validate
	if err := validator.Struct(cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}
