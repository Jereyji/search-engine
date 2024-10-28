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

type WordList struct {
	db *pgxpool.Pool
	mu sync.RWMutex
}

func NewWordList(db *pgxpool.Pool) *WordList {
	return &WordList{db: db}
}

func (s *WordList) Create(context context.Context, wordList *entity.WordList) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var id int
	err := s.db.QueryRow(context, queries.AddWordList, wordList.Word, wordList.IsFiltred).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *WordList) Word(context context.Context, id int) (*entity.WordList, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var wordList entity.WordList
	err := s.db.QueryRow(context, queries.GetWordList, id).Scan(&wordList.ID, &wordList.Word, &wordList.IsFiltred)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &wordList, nil
		}

		return nil, err
	}

	return &wordList, nil
}

func (s *WordList) Update(context context.Context, wordList *entity.WordList) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(context, queries.UpdateWordList, wordList.Word, wordList.IsFiltred, wordList.ID)
	return err
}

func (s *WordList) Delete(context context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(context, queries.DeleteWordList, id)
	return err
}
