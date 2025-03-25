package config

import (
	"github.com/spf13/viper"
	"log"
)

type App struct {
	Name     string
	Port     string
	Version  string
	LogLevel string
	Postgres Postgres
}

type Postgres struct {
	Host        string
	Port        string
	Username    string
	Password    string
	DBName      string
	SSLMode     string
	MaxIdleTime int
	MaxLifeTime int
	MaxOpenConn int
	MaxIdleConn int
}

func ConfigInit() (*App, error) {
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv() // Mengambil nilai dari ENV jika tersedia

	// Membaca file .env
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No .env file found, using environment variables: %v", err)
		return nil, err
	}

	app := &App{
		Name:     viper.GetString("APP_NAME"),
		Port:     viper.GetString("APP_PORT"),
		Version:  viper.GetString("APP_VERSION"),
		LogLevel: viper.GetString("APP_LOG_LEVEL"),
		Postgres: Postgres{
			Host:        viper.GetString("POSTGRES_HOST"),
			Port:        viper.GetString("POSTGRES_PORT"),
			Username:    viper.GetString("POSTGRES_USER"),
			Password:    viper.GetString("POSTGRES_PASSWORD"),
			DBName:      viper.GetString("POSTGRES_DBNAME"),
			SSLMode:     viper.GetString("POSTGRES_SSL_MODE"),
			MaxIdleTime: viper.GetInt("POSTGRES_MAX_IDLE_TIME"),
			MaxLifeTime: viper.GetInt("POSTGRES_MAX_LIFE_TIME"),
			MaxOpenConn: viper.GetInt("POSTGRES_MAX_OPEN_CONN"),
			MaxIdleConn: viper.GetInt("POSTGRES_MAX_IDLE_CONN"),
		},
	}

	return app, nil

}
