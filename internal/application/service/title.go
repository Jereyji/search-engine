package service

import (
	"context"
	"errors"

	"github.com/Jereyji/search-engine/internal/domain/entity"
)

type TitleInfo struct {
	Link     string
	LinkText string
}

func (p *TitleInfo) storeToRepository(ctx context.Context, s *CrawlerService, URL *entity.URLList) (*Response, error) {
	res := Response{
		URL: URL.Link,
	}

	relatedURL, err := s.ensureURLExists(ctx, p.Link)
	if err != nil {
		return nil, err
	}

	if errors.Is(err, errors.New(errExistRow)) {
		return &res, nil
	}

	curCountWords, curCountFilteredWords, err := s.addText(ctx, p.LinkText, relatedURL.ID, 0)
	if err != nil {
		return nil, err
	}

	res.CountWords += curCountWords
	res.CountFilteredWords += curCountFilteredWords

	return &res, nil
}
