package config

import (
	"log"

	"github.com/Kushkaftar/leadCannon/internal/models"
	"github.com/spf13/viper"
)

var config models.Config

func NewConfig() (*models.Config, error) {

	if err := initConfig(); err != nil {
		log.Fatalf("crush init config, %v", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &config, nil
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	return viper.ReadInConfig()
}
