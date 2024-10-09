package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Region string
	Tags   []string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}
	AppConfig.Region = viper.GetString("aws.region")
	AppConfig.Tags = viper.GetStringSlice("tags")
}
