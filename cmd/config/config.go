package config

import (
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	DBName string `mapstructure:"database"`
	DBUSer string `mapstructure:"user"`
	DBPass string `mapstructure:"password"`
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
