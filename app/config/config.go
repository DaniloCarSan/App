package config

import (
	"github.com/spf13/viper"
)

var configInstance *config

type config struct {
	API      APIConfig
	Database DBConfig
}

func init() {

	err := Load()

	if err != nil {
		panic(err)
	}
}

func Load() error {

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	configInstance = &config{
		API: APIConfig{
			Host: viper.GetString("API_HOST"),
			Port: viper.GetString("API_PORT"),
		},
		Database: DBConfig{
			Host: viper.GetString("DB_HOST"),
			Port: viper.GetString("DB_PORT"),
			User: viper.GetString("DB_USER"),
			Pass: viper.GetString("DB_PASS"),
			Name: viper.GetString("DB_NAME"),
		},
	}

	return nil
}

func GetConfig() *config {
	return configInstance
}

func GetDB() *DBConfig {
	return &configInstance.Database
}

func GetAPI() *APIConfig {
	return &configInstance.API
}
