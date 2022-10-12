package config

import "github.com/spf13/viper"

const DEFAULT_DB_HOST = "127.0.0.1"
const DEFAULT_DB_PORT = "5432"
const DEFAULT_DB_NAME = "postgres"
const DEFAULT_DB_USER = "postgres"
const DEFAULT_DB_PASS = "postgres"

type DBConfig struct {
	Host string
	Port string
	Name string
	User string
	Pass string
}

func init() {

	viper.SetDefault("DB_HOST", DEFAULT_DB_HOST)
	viper.SetDefault("DB_PORT", DEFAULT_DB_PORT)
	viper.SetDefault("DB_NAME", DEFAULT_DB_NAME)
	viper.SetDefault("DB_USER", DEFAULT_DB_USER)
	viper.SetDefault("DB_PASS", DEFAULT_DB_PASS)

}
