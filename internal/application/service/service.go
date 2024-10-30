package service

import (
	"github.com/Jereyji/search-engine/internal/domain/repository"
)

type CrawlerService struct {
	repository repository_interface.CrawlerRepositoryInterface
}

func NewCrawlerService(repository repository_interface.CrawlerRepositoryInterface) *CrawlerService {
	return &CrawlerService{
		repository: repository,
	}
}

type Response struct {
	URL                string
	CountWords         int
	CountFilteredWords int
}
