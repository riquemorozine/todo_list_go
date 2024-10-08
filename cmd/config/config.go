package config

import (
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	DBName       string `mapstructure:"DB_NAME"`
	DBUSer       string `mapstructure:"DB_USER"`
	DBPass       string `mapstructure:"DB_PASS"`
	JWTSecret    string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth    string
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
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
