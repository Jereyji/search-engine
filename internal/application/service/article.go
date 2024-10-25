package service

import (
	"context"

	"github.com/Jereyji/search-engine/internal/domain/entity"
)

type ArticleInfo struct {
	MainText         string
	RelatedLinks     []string
	RelatedLinkTexts []string
}

func (p *ArticleInfo) storeToRepository(ctx context.Context, s *CrawlerService, URL *entity.URLList) (*Response, error) {
	res := Response{
		URL: URL.Link,
	}

	if p.MainText != "" {
		curCountWords, curCountFilteredWords, err := s.addText(ctx, p.MainText, URL.ID, 0)
		if err != nil {
			return nil, err
		}

		res.CountWords += curCountWords
		res.CountFilteredWords += curCountFilteredWords
	}

	for i, relatedLink := range p.RelatedLinks {
		relatedURL, err := s.ensureURLExists(ctx, relatedLink)
		if err != nil {
			return nil, err
		}

		linkBetweenID, err := s.repository.AddLinkBetweenURLs(ctx, &entity.LinkBetweenURL{FromURLID: URL.ID, ToURLID: relatedURL.ID})
		if err != nil {
			return nil, err
		}

		curCountWords, curCountFilteredWords, err := s.addText(ctx, p.RelatedLinkTexts[i], relatedURL.ID, linkBetweenID)
		if err != nil {
			return nil, err
		}

		res.CountWords += curCountWords
		res.CountFilteredWords += curCountFilteredWords
	}

	return &res, nil
}
