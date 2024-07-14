package config

import (
	"github.com/spf13/viper"
	"log"
)

func NewViper() *viper.Viper {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("./../")
	config.AddConfigPath("./")
	if err := config.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	return config
}
