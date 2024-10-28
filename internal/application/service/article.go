package service

import (
	"context"

	"github.com/Jereyji/search-engine/internal/domain/entity"
)

type ArticleInfo struct {
	MainText         string
	RelatedLinks     []TitleInfo
}

func (p *ArticleInfo) storeToRepository(ctx context.Context, s *CrawlerService, URL *entity.URLList) ([]Response, error) {
	res := []Response{{URL: URL.Link}}

	if p.MainText != "" {
		curCountWords, curCountFilteredWords, err := s.addText(ctx, &p.MainText, URL.ID, 0)
		if err != nil {
			return nil, err
		}

		res[0].CountWords += curCountWords
		res[0].CountFilteredWords += curCountFilteredWords
	}

	for _, relatedLink := range p.RelatedLinks {
		curRes, err := relatedLink.storeToRepository(ctx, s, URL)
		if err != nil {
			return nil, err
		}
		
		res = append(res,  *curRes)
	}

	return res, nil
}
