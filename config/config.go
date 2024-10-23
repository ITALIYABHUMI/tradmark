package config

import (
	"log"

	"github.com/spf13/viper"

	"github.com/tradmark/public/model"
)

var settings model.Settings

func Init() {
	err := LoadConfig()
	if err != nil {
		panic(err)
	}

	Connectdb()
	EsClientConnection()
	EsCreateIndexIfNotExists()
}

func GetConfig() *model.Settings {
	return &settings
}

func LoadConfig() (err error) {
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v\n", err)
		return err
	}

	if err = viper.Unmarshal(&settings); err != nil {
		log.Printf("Error unmarshalling config: %v\n", err)
		return err
	}

	return nil
}
