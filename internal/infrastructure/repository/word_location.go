package repository

import (
	"context"
	"errors"

	"github.com/Jereyji/search-engine/internal/domain/entity"
	"github.com/Jereyji/search-engine/internal/infrastructure/repository/queries"
	"github.com/jackc/pgx/v5"
)

func (s *CrawlerRepository) CreateWordLocation(context context.Context, wordLocation *entity.WordLocation) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var id int
	err := s.db.QueryRow(context, queries.AddWordLocation, wordLocation.WordID, wordLocation.URLID, wordLocation.Location).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *CrawlerRepository) WordLocation(context context.Context, id int) (*entity.WordLocation, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var wordLocation entity.WordLocation
	err := s.db.QueryRow(context, queries.GetWordLocation, id).Scan(&wordLocation.ID, &wordLocation.WordID, &wordLocation.URLID, &wordLocation.Location)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &wordLocation, nil
		}

		return nil, err
	}

	return &wordLocation, nil
}

func (s *CrawlerRepository) UpdateWordLocation(context context.Context, wordLocation *entity.WordLocation) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(context, queries.UpdateWordLocation, wordLocation.WordID, wordLocation.URLID, wordLocation.Location, wordLocation.ID)
	return err
}

func (s *CrawlerRepository) DeleteWordLocation(context context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(context, queries.DeleteWordLocation, id)
	return err
}
