package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DataLinks []struct {
		Url      string `yaml:"url"`
		Selector string `yaml:"selector"`
		Text     string `yaml:"text"`
	} `yaml:"data_links"`

	DataBaseURL string `yaml:"database_url"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, err
	}

	// if err := cleanenv.ReadEnv(&cfg); err != nil {
	// 	return nil, err
	// }

	return &cfg, nil
}
