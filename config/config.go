package config

import (
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

var cfg Config

func Parse() *Config {
	errGodot := godotenv.Load()
	if errGodot != nil {
		panic("Error loading .env file")
	}
	err := config.FromEnv().To(&cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
