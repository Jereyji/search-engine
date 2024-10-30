package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

// Определяем структуру для конфигурации
type Config struct {
	DataURLs []struct {
		URL            string `yaml:"url"`
		TitleTextTag   string `yaml:"title_text_tag"`
		TitleLinkTag   string `yaml:"title_link_tag"`
		ArticleTextTag string `yaml:"article_text_tag"`
		ArticleLinkTag string `yaml:"article_link_tag"`
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