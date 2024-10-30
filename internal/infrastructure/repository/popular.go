package repository

import (
	"context"

	"github.com/Jereyji/search-engine/internal/domain/entity"
	"github.com/Jereyji/search-engine/internal/infrastructure/repository/queries"
)

func (s *CrawlerRepository) PopularDomains(ctx context.Context) ([]entity.PopularDomain, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	rows, err := s.db.Query(ctx, queries.MostIndexedDomains)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var domains []entity.PopularDomain
	for rows.Next() {
		var domain entity.PopularDomain
		if err := rows.Scan(&domain.Domain, &domain.Count); err != nil {
			return nil, err
		}
		domains = append(domains, domain)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return domains, nil
}

func (s *CrawlerRepository) PopularWords(ctx context.Context) ([]entity.PopularWord, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	rows, err := s.db.Query(ctx, queries.PopularWords)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var words []entity.PopularWord
	for rows.Next() {
		var word entity.PopularWord
		if err := rows.Scan(&word.Word, &word.Count); err != nil {
			return nil, err
		}
		words = append(words, word)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return words, nil
}
