package service

import (
	"context"

	"github.com/Jereyji/search-engine/internal/domain/entity"
)

func (s *CrawlerService) PopularDomains(ctx context.Context) ([]entity.PopularDomain, error) {
	res, err := s.repository.PopularDomains(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CrawlerService) PopularWords(ctx context.Context) ([]entity.PopularWord, error) {
	res, err := s.repository.PopularWords(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}