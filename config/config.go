package config

import "os"

type config struct {
	DBUsername string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	ServerPort string
}

var Config *config

func init() {
	Config = &config{
		DBUsername: os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBName:     os.Getenv("POSTGRES_DB"),
		DBHost:     os.Getenv("POSTGRES_HOST"),
		DBPort:     os.Getenv("POSTGRES_PORT"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}
