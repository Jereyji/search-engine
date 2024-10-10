package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DataBase struct {
		Username string `yaml:"username" env:"DB_USERNAME"`
		Password string `yaml:"password" env:"DB_PASSWORD"`
		Host     string `yaml:"host" env:"DB_HOST"`
		Port     string `yaml:"port" env:"DB_PORT"`
		Dbname   string `yaml:"dbname" env:"DB_NAME"`
		Sslmode  string `yaml:"sslmode" env:"DB_SSLMODE"`
	} `yaml:"db"`
	
	Links []string `yaml:"links"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
