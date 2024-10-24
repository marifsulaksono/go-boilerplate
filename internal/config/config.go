package config

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Configuration struct {
	App      App      `json:"app"`
	Database Database `json:"database"`
	JWT      JWT      `json:"jwt"`
}

var Config *Configuration

func Load(ctx context.Context, isEnvFile bool) error {
	// load env
	if isEnvFile {
		if err := loadEnvFile(); err != nil {
			return err
		}
	} else {
		filename := "config"
		ext := "yaml"
		path := "./config"
		if err := loadConfigFile(filename, ext, path); err != nil {
			return err
		}
	}

	// prepare configuration values
	Config = &Configuration{
		App: App{
			Port: viper.GetInt("app.port"),
		},
		Database: Database{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			Username: viper.GetString("database.username"),
			Password: viper.GetString("database.password"),
			Name:     viper.GetString("database.name"),
		},
		JWT: JWT{
			AccessSecret:       viper.GetString("jwt.access_secret_key"),
			RefreshSecret:      viper.GetString("jwt.refresh_secret_key"),
			AccessExpiryInSec:  viper.GetInt("jwt.access_expiry_in_second"),
			RefreshExpiryInSec: viper.GetInt("jwt.refresh_expiry_in_second"),
		},
	}

	return nil
}

// use this function if using file .env
func loadEnvFile() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	viper.AutomaticEnv()

	return nil
}

// use this function if not using file .env
func loadConfigFile(filename, ext, path string) error {
	viper.SetConfigName(filename)
	viper.SetConfigType(ext)
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
