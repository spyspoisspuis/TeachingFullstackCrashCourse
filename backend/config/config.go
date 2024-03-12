package config

import (
	"log"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App AppConfig
		Db DbConfig
	}

	AppConfig struct {
		Port int
	}

	DbConfig struct {
		Host     string
		User     string
		Password string
		Dbname   string
		Port     int
		Sslmode  string
		Timezone string
	}
)

func InitializeViper(path string) {
	log.Println("Viper initializing....")
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	log.Println("Viper initialized!")
}

func GetConfig() Config {
	log.Println("Getting config....")

	return Config {
		Db: DbConfig {
			Host:     viper.GetString("DB_HOST"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			Dbname:   viper.GetString("DB_NAME"),
			Port:     viper.GetInt("DB_PORT"),
			Sslmode:  viper.GetString("DB_SSLMODE"),
			Timezone: viper.GetString("DB_TIMEZONE"),
		},
		App: AppConfig {
			Port: viper.GetInt("PORT"),
		},
	}
}