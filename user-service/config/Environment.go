package config

import (
	"github.com/spf13/viper"
	"log"
)

func LoadEnvironment() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to load env: %v", err)
	}
}
