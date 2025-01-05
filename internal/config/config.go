// Package config has parser configuration of application bloggo.
package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config - object of configuration.
type Config struct {
	HTTP    HTTP    `yaml:"http"`
	Storage Storage `yaml:"storage"`
	Logger  Logger  `yaml:"logger"`
}

// HTTP - has fields for server.
type HTTP struct {
	Port        string        `yaml:"port"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

// Storage - has connection for database.
type Storage struct {
	Path string `yaml:"path"`
}

// Logger - has envierenment of work application.
type Logger struct {
	Env string `yaml:"env"`
}

// New - create new configuration.
func New() *Config {
	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		log.Fatal("CONFIG_PATH is empty. Please, add path to config in .env")
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exists: %s", cfgPath)
	}

	var cfg Config
	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		log.Fatalf("cannot read config file: %s", err)
	}

	return &cfg
}
