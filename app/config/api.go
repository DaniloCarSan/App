package config

import "github.com/spf13/viper"

const DEFAULT_API_HOST = "127.0.0.1"
const DEFAULT_API_PORT = "9000"

type APIConfig struct {
	Host string
	Port string
}

func init() {
	viper.SetDefault("API_HOST", DEFAULT_API_HOST)
	viper.SetDefault("API_PORT", DEFAULT_API_PORT)
}
