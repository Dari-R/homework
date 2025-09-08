package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	key := os.Getenv("KEY")
	if key == "" {
		panic("Не передан параметр Key в переменную окружения")
	}

	return &Config{
		Key: key,
	}
}
