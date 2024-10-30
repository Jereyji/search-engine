package handler

import (
	"github.com/Jereyji/search-engine/internal/application/service"
	"github.com/Jereyji/search-engine/internal/pkg/config"
)

type CrawlerHandler struct {
	crawlerService *service.CrawlerService
	config         *config.Config
}

func NewCrawlerHandler(crawlerService *service.CrawlerService, cfg *config.Config) *CrawlerHandler {
	return &CrawlerHandler{
		crawlerService: crawlerService,
		config:         cfg,
	}
}
