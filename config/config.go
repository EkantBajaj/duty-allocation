package config

import (
	"log"

	"github.com/spf13/viper"
)

// LoadConfig loads the configuration from the config.yaml file
func LoadConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return err
	}

	// Additional configuration setup if needed

	return nil
}
