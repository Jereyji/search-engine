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

type LinkBetweenURLs struct {
	db *pgxpool.Pool
	mu sync.RWMutex
}

func NewLinkBetweenURLs(db *pgxpool.Pool) *LinkBetweenURLs {
	return &LinkBetweenURLs{db: db}
}

func (s *LinkBetweenURLs) Create(context context.Context, linkBetweenURL *entity.LinkBetweenURL) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var id int
	err := s.db.QueryRow(context, queries.AddLinkBetweenURL, linkBetweenURL.FromURLID, linkBetweenURL.ToURLID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *LinkBetweenURLs) LinkBetweenURLs(context context.Context, id int) (*entity.LinkBetweenURL, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var linkBetweenURL entity.LinkBetweenURL
	err := s.db.QueryRow(context, queries.GetLinkBetweenURL, id).Scan(&linkBetweenURL.ID, &linkBetweenURL.FromURLID, &linkBetweenURL.ToURLID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &linkBetweenURL, nil
		}

		return nil, err
	}

	return &linkBetweenURL, nil
}

func (s *LinkBetweenURLs) Update(context context.Context, linkBetweenURL *entity.LinkBetweenURL) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(context, queries.UpdateLinkBetweenURL, linkBetweenURL.FromURLID, linkBetweenURL.ToURLID, linkBetweenURL.ID)
	return err
}

func (s *LinkBetweenURLs) Delete(context context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(context, queries.DeleteLinkBetweenURL, id)
	return err
}
