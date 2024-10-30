package service

import (
	"context"

	"github.com/Jereyji/search-engine/internal/domain/entity"
)

type TitleInfo struct {
	Link     string
	LinkText string
}

func (p *TitleInfo) storeToRepository(ctx context.Context, s *CrawlerService, URL *entity.URLList) (*Response, error) {
	relatedURL, err := s.ensureURLExists(ctx, p.Link)
	if err != nil {
		return nil, err
	}

	if relatedURL.Is_parsed {
		return nil, nil
	}

	linkBetweenID, err := s.repository.CreateLinkBetweenURLs(ctx, &entity.LinkBetweenURL{FromURLID: URL.ID, ToURLID: relatedURL.ID})
	if err != nil {
		return nil, err
	}

	curCountWords, curCountFilteredWords, err := s.addText(ctx, &p.LinkText, relatedURL.ID, linkBetweenID)
	if err != nil {
		return nil, err
	}

	res := Response{
		URL:                relatedURL.Link,
		CountWords:         curCountWords,
		CountFilteredWords: curCountFilteredWords,
	}

	return &res, nil
}
