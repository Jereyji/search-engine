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

type LinkWord struct {
	db *pgxpool.Pool
	mu sync.RWMutex
}

func NewLinkWord(db *pgxpool.Pool) *LinkWord {
	return &LinkWord{db: db}
}

func (s *LinkWord) Create(context context.Context, linkWord *entity.LinkWord) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var id int
	err := s.db.QueryRow(context, queries.AddLinkWord, linkWord.WordID, linkWord.LinkID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *LinkWord) LinkWord(context context.Context, id int) (*entity.LinkWord, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var linkWord entity.LinkWord
	err := s.db.QueryRow(context, queries.GetLinkWord, id).Scan(&linkWord)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &linkWord, nil
		}

		return nil, err
	}

	return &linkWord, nil
}

func (s *LinkWord) Update(context context.Context, linkWord *entity.LinkWord) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(context, queries.UpdateLinkWord, linkWord.WordID, linkWord.LinkID, linkWord.ID)
	return err
}

func (s *LinkWord) Delete(context context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(context, queries.DeleteLinkWord, id)
	return err
}
