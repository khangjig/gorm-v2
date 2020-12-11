package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

var config *Config

type Config struct {
	Port  string `envconfig:"PROJECT_PORT"`
	Debug bool   `envconfig:"IS_DEBUG"`

	MySQL struct {
		Host   string `envconfig:"DB_HOST"`
		Port   string `envconfig:"DB_PORT"`
		DBName string `envconfig:"DB_NAME"`
		User   string `envconfig:"DB_USER"`
		Pass   string `envconfig:"DB_PASS"`
	}
}

func init() {
	config = &Config{}

	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load godotenv", err)
	}

	err = envconfig.Process("", config)
	if err != nil {
		err = errors.Wrap(err, "Failed to decode config env")
		log.Fatal(err)
	}

	if len(config.Port) == 0 {
		config.Port = "6000"
	}
}

func GetConfig() *Config {
	return config
}
