package config

import (
	"log"

	"github.com/spf13/viper"
	
	"github.com/tradmark/model"
)

var settings model.Settings

func Init() {
	err := LoadConfig()
	if err != nil {
		panic(err)
	}

	Connectdb()
}

func GetConfig() *model.Settings {
	return &settings
}

func LoadConfig() (err error) {
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file: %v\n", err)
		return err
	}

	err = viper.Unmarshal(&settings)

	return nil
}
