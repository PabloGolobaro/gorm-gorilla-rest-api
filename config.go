package main

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port             string `mapstructure:"port"`
	ConnectionString string `mapstructure:"connection_string"`
}

var AppConfig *Config

func LoadAppConfig() {
	log.Println("Loading server Configuration...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
