package config

import (
	"fmt"

	"github.com/JeremyLoy/config"
	"github.com/joho/godotenv"
)

type Config struct {
	Debug      string
	HTTPServer string
	HTTPPort   string
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPass     string
	DBSchema   string
}

func Parse() (cfg Config) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	err = config.FromEnv().To(&cfg)
	if err != nil {
		fmt.Println(err)
	}

	return cfg
}
