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
	Redis    Redis    `json:"redis"`
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

	/*
		viper get value is depend on configuration file.
		if using file .env, please use variable name. e.g APP_PORT
		if using file yaml, toml, or file with tree concept, please use path name, e.g app.port

		more info contact me @marifsulaksono
	*/

	// prepare configuration values
	Config = &Configuration{
		App: App{
			Port: viper.GetInt("APP_PORT"),
		},
		Database: Database{
			Host:     viper.GetString("DATABASE_HOST"),
			Port:     viper.GetString("DATABASE_PORT"),
			Username: viper.GetString("DATABASE_USERNAME"),
			Password: viper.GetString("DATABASE_PASSWORD"),
			Name:     viper.GetString("DATABASE_NAME"),
		},
		JWT: JWT{
			PrivateKeyPathFile: viper.GetString("JWT_PRIVATE_KEY_PATH_FILE"),
			PublicKeyPathFile:  viper.GetString("JWT_PUBLIC_KEY_PATH_FILE"),
			AccessSecret:       viper.GetString("JWT_ACCESS_SECRET_KEY"),
			RefreshSecret:      viper.GetString("JWT_REFRESH_SECRET_KEY"),
			AccessExpiryInSec:  viper.GetInt("JWT_ACCESS_EXPIRY_IN_SECOND"),
			RefreshExpiryInSec: viper.GetInt("JWT_REFRESH_EXPIRY_IN_SECOND"),
		},
		Redis: Redis{
			Host:     viper.GetString("REDIS_HOST"),
			Port:     viper.GetString("REDIS_PORT"),
			Password: viper.GetString("REDIS_PASSWORD"),
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
