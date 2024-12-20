package config

import "github.com/spf13/viper"

type Config struct {
	DBUrl string `mapstructure:"DB_URL"`
}

func LoadEnv(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
