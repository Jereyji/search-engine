package repository

import (
	"context"
	"errors"

	"github.com/Jereyji/search-engine/internal/domain/entity"
	"github.com/Jereyji/search-engine/internal/infrastructure/repository/queries"
	"github.com/jackc/pgx/v5"
)

func (s *CrawlerRepository) CreateURL(context context.Context, url *entity.URLList) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var id int
	err := s.db.QueryRow(context, queries.AddUrlList, url.Link, url.Is_parsed).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *CrawlerRepository) URL(context context.Context, link string) (*entity.URLList, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var url entity.URLList
	err := s.db.QueryRow(context, queries.GetUrlList, link).Scan(&url.ID, &url.Link)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &url, nil
		}

		return nil, err
	}

	return &url, nil
}

func (s *CrawlerRepository) UpdateURL(context context.Context, url *entity.URLList) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(context, queries.UpdateUrlList, url.Link, url.Is_parsed, url.ID)
	return err
}

func (s *CrawlerRepository) DeleteURL(context context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(context, queries.DeleteUrlList, id)
	return err
}
