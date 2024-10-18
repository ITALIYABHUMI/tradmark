package config

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/tradmark/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	DB *gorm.DB
}

var Database DbInstance
var gormDB *gorm.DB

func Connectdb() {

	cfg := model.DatabaseConfig{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		DBName:   viper.GetString("DB_NAME"),
		SSLMode:  viper.GetString("DB_SSL_MODE"),
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to the database successfully")
	Database = DbInstance{DB: db}

	err = db.AutoMigrate(&model.CaseFile{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return Database.DB
}

func GetDbWithContext(ctx context.Context) *gorm.DB {
	return GetDB().WithContext(ctx)
}
