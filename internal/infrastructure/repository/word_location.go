package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/Jereyji/search-engine/internal/domain/entity"
	"github.com/Jereyji/search-engine/internal/infrastructure/repository/queries"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WordLocation struct {
	db *pgxpool.Pool
	mu sync.RWMutex
}

func NewWordLocation(db *pgxpool.Pool) *WordLocation {
	return &WordLocation{db: db}
}

func (s *WordLocation) Create(context context.Context, wordLocation *entity.WordLocation) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var id int
	err := s.db.QueryRow(context, queries.AddWordLocation, wordLocation.WordID, wordLocation.URLID, wordLocation.Location).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *WordLocation) WordLocation(context context.Context, id int) (*entity.WordLocation, error) {
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

func (s *WordLocation) Update(context context.Context, wordLocation *entity.WordLocation) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(context, queries.UpdateWordLocation, wordLocation.WordID, wordLocation.URLID, wordLocation.Location, wordLocation.ID)
	return err
}

func (s *WordLocation) Delete(context context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(context, queries.DeleteWordLocation, id)
	return err
}
