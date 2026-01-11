package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Port            string `env:"PORT"`
	RedisHost       string `env:"REDIS_HOST"`
	RedisPort       string `env:"REDIS_PORT"`
	RedisDB         int    `env:"REDIS_DB"`
	RedisPass       string `env:"REDIS_PASSWORD"`
	RateLimit       int    `env:"RATE_LIMIT"`
	RateLimitWindow int    `env:"RATE_LIMIT_WINDOW"`
	RedisAddress    string `env:"REDIS_ADDRESS"`
}

func GetConfig() Configuration {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	configuration := Configuration{}

	err = gonfig.GetConf("", &configuration)
	if err != nil {
		fmt.Println("Error loading config:", err)
	}

	return configuration
}
