package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

// Определяем структуру для конфигурации
type Config struct {
	DataURLs []struct {
		URL             string `yaml:"url"`
		TitleSelector   string `yaml:"title_selector"`
		TitleText       string `yaml:"title_text"`
		ArticleSelector string `yaml:"article_selector"`
		ArticleLink     string `yaml:"article_link"`
	} `yaml:"data_links"`
}

// Функция для загрузки конфигурации
func LoadConfig(path string) (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
